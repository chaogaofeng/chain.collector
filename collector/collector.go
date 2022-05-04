package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/glodnet/chain.collector/schema"
	"github.com/glodnet/chain.go/restclient"
	"github.com/glodnet/chain/app"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
)

type collector struct {
	logger log.Logger
	db     *gorm.DB
	client *restclient.RestClient
	cdc    codec.Codec
}

func NewCollector(logger log.Logger, db *gorm.DB, client *restclient.RestClient) *collector {
	encodingConfig := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	return &collector{
		logger: logger,
		db:     db,
		client: client,
		cdc:    encodingConfig.Marshaler,
	}
}

// Start starts to synchronize Binance Chain data.
func (cl *collector) Start(ctx context.Context) {
	start := cl.GetLastBlockHeight() + 1
	cl.client.Collect(ctx, start, cl)
}

func hash(bytes []byte) string {
	return fmt.Sprintf("%X", tmhash.Sum(bytes))
}

func (cl *collector) GetLastBlockHeight() int64 {
	var mBlock schema.Block
	if err := cl.db.Order("height desc").Find(&mBlock).Error; err != nil {
		panic(err)
	}
	return mBlock.Height
}

func (cl *collector) Logger() log.Logger {
	return cl.logger
}

func (cl *collector) HandleGenesis(genesisState map[string]json.RawMessage) error {
	return nil
}

func (cl *collector) HandlePrevBlock(block *tmservice.GetBlockByHeightResponse) error {
	cl.logger.Debug("HandlePrevBlock", "height", block.Block.Header.Height)
	var mBlock schema.Block
	if err := cl.db.Last(&mBlock).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	prevHash := hash(block.Block.Header.LastBlockId.Hash)
	hash := hash(block.BlockId.Hash)
	height := block.Block.Header.Height
	if mBlock.BlockHash != hash && mBlock.ParentHash != prevHash && mBlock.Height != height {
		return fmt.Errorf("database mismatch error")
	}
	return nil
}
func (cl *collector) HandleBlock(block *tmservice.GetBlockByHeightResponse, txs []*tx.GetTxResponse) error {
	cl.logger.Debug("HandleBlock", "height", block.Block.Header.Height)

	// block
	mblock := &schema.Block{
		Height:        block.Block.Header.Height,
		Proposer:      sdk.AccAddress(block.Block.Header.ProposerAddress).String(),
		BlockHash:     hash(block.BlockId.Hash),
		ParentHash:    hash(block.Block.Header.LastBlockId.Hash),
		NumPrecommits: int64(len(block.Block.LastCommit.Signatures)),
		NumTxs:        int64(len(block.Block.Data.Txs)),
		Timestamp:     block.Block.Header.Time,
	}
	mblock.Moniker = schema.QueryValidatorMoniker(cl.db, mblock.Proposer)

	// block LastCommit
	precommits := make([]*schema.PreCommit, mblock.NumPrecommits, mblock.NumPrecommits)
	if mblock.NumPrecommits > 0 {
		valSets, err := cl.client.ValidatorSetByHeight(block.Block.LastCommit.Height, nil)
		if err != nil {
			return fmt.Errorf("failed to query validator by height %d: %s", block.Block.LastCommit.Height, err)
		}
		valMaps := map[string]*tmservice.Validator{}
		for _, validator := range valSets.Validators {
			valMaps[validator.Address] = validator
		}
		for i, precommit := range block.Block.LastCommit.Signatures {
			addr := sdk.ConsAddress(precommit.ValidatorAddress).String()
			val := valMaps[addr]
			pc := &schema.PreCommit{
				Height:           block.Block.LastCommit.Height,
				Round:            block.Block.LastCommit.Round,
				ValidatorAddress: addr,
				VotingPower:      val.VotingPower,
				ProposerPriority: val.ProposerPriority,
				Timestamp:        precommit.Timestamp,
			}
			precommits[i] = pc
		}
	}

	// block txs
	mtxs := make([]*schema.Transaction, mblock.NumTxs, mblock.NumTxs)
	for i, tx := range txs {
		t := &schema.Transaction{
			Height:    tx.TxResponse.Height,
			TxHash:    tx.TxResponse.TxHash,
			Code:      tx.TxResponse.Code,
			RawLog:    tx.TxResponse.RawLog,
			Memo:      tx.Tx.Body.Memo,
			Fees:      tx.Tx.GetFee().String(),
			GasWanted: tx.TxResponse.GasWanted,
			GasUsed:   tx.TxResponse.GasUsed,
			Timestamp: block.Block.Header.Time,
		}
		mtxs[i] = t
	}

	return cl.db.Transaction(func(db *gorm.DB) error {
		return nil
	})
	return nil
}
