package cron

import (
	"github.com/robfig/cron"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
)

// Cron wraps all required parameters to create cron jobs
type Cron struct {
	logger log.Logger
	db     *gorm.DB
}

// NewCron sets necessary config and clients to begin jobs
func NewCron(logger log.Logger, db *gorm.DB) *Cron {
	return &Cron{
		logger: logger,
		db:     db,
	}
}

// Start starts to create cron jobs which fetches chosen asset list information and
// store them in database every hour and every 24 hours.
func (c *Cron) Start() {
	c.logger.Info("Starting cron ...")

	cron := cron.New()

	// Every hour
	cron.AddFunc("0 0 * * * *", func() {
	})

	// Every 24 hours at 0:00 AM UTC timezone
	cron.AddFunc("0 0 0 * * *", func() {

	})

	cron.Start()
}
