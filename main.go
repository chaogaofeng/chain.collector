package main

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain.collector/collector"
	"github.com/glodnet/chain.collector/config"
	"github.com/glodnet/chain.collector/db"
	"github.com/glodnet/chain.collector/logger"
	"github.com/glodnet/chain.go/restclient"
)

func main() {
	// Parse config from configuration file (config.yaml).
	config := config.ParseConfig()

	// Create new logger with log configruation.
	logger := logger.NewLogger(config.Log.Level, config.Log.Dir)

	// Create new db with database configruation.
	db := db.MustNewDB(config.DB.Mode, config.DB.DSN)

	// Create new client with node configruation.
	client := restclient.New(config.Node.APIEndpoint, "", sdk.DecCoin{}, sdk.Dec{}, config.Node.AddressPrefix)

	ctx := context.Background()

	// Start collect chain data.
	collector := collector.NewCollector(logger, db, client)
	collector.Start(ctx)
}
