package schema

import (
	"gorm.io/gorm"
	"time"
)

// Block defines the structure for block information.
type Block struct {
	ID                   uint      `json:"-" gorm:"primarykey" `
	Height               int64     `json:"block_height" gorm:"not null"`
	Proposer             string    `json:"propopser_addr" gorm:"not null"`
	Moniker              string    `json:"propopser_moniker" gorm:"not null"`
	BlockHash            string    `json:"block_hash" gorm:"not null,unique"`
	ParentHash           string    `json:"parent_hash" gorm:"not null"`
	PrecommitNum         int64     `json:"precommit_validator_num" gorm:"not null"`
	PrecommitVotingPower int64     `json:"precommit_voting_power"`
	TotalValidatorNum    int64     `json:"total_validator_num"`
	TotalVotingPower     int64     `json:"total_voting_power"`
	NumTxs               int64     `json:"num_txs" gorm:"default:0"`
	Timestamp            time.Time `json:"timestamp"`
}

func QueryBlockLatestHeight(db *gorm.DB) (*Block, error) {
	var item Block
	if res := db.Order("height desc").Find(&item); res.Error != nil {
		return nil, res.Error
	} else if res.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}

func QueryBlockByHeight(db *gorm.DB, height int64) (*Block, error) {
	var item Block
	if res := db.Model(&Block{}).Where("height = ?", height).Find(&item); res.Error != nil {
		return nil, res.Error
	} else if res.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}

func QueryBlocks(db *gorm.DB, page int, size int) ([]*Block, int64, error) {
	offset := (page - 1) * size

	var items []*Block
	var total int64
	if res := db.Model(&Block{}).Order("height desc").Offset(offset).Limit(size).Find(&items).Count(&total); res.Error != nil {
		return nil, 0, res.Error
	} else if res.RowsAffected == 0 {
		return nil, 0, nil
	}
	return items, 0, nil
}
