package main

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain.collector/collector"
	"github.com/glodnet/chain.collector/config"
	"github.com/glodnet/chain.collector/cron"
	"github.com/glodnet/chain.collector/db"
	"github.com/glodnet/chain.collector/logger"
	"github.com/glodnet/chain.collector/rest"
	"github.com/glodnet/chain.go/restclient"
	"log"
	"os"
	"os/signal"
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

	ctx, cancel := context.WithCancel(context.Background())

	// Start collect chain data.
	collector := collector.NewCollector(logger, db, client)
	go collector.Start(ctx)

	// Start cron chain data.
	cron := cron.NewCron(logger, db)
	go cron.Start()

	// Start the API server
	api := rest.NewApiServer(logger, db, client)
	go api.Start(config.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received.
	sig := <-c
	cancel()

	log.Println("shutting down the server: ", sig)
}
