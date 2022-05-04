package schema

import (
	"gorm.io/gorm"
	"time"
)

// PreCommit defines the schema for precommit state information.
type PreCommit struct {
	ID               uint      `json:"-" gorm:"primarykey" `
	Height           int64     `json:"height" grom:"not null"`
	Round            int32     `json:"round" grom:"not null"`
	ConsensusAddress string    `json:"consensus"`
	IsProposer       bool      `json:"is_proposer"`
	ValidatorAddress string    `json:"validator_address" grom:"not null"`
	VotingPower      int64     `json:"voting_power" grom:"not null"`
	ProposerPriority int64     `json:"proposer_priority" grom:"not null"`
	Timestamp        time.Time `json:"timestamp"`
}

func QueryPreCommitsByHeight(db *gorm.DB, height int64, page int, size int) ([]*PreCommit, int64, error) {
	offset := (page - 1) * size

	var items []*PreCommit
	var total int64
	if res := db.Model(&PreCommit{}).Where("height = ?", height).Offset(offset).Limit(size).Find(&items).Count(&total); res.Error != nil {
		return nil, 0, res.Error
	} else if res.RowsAffected == 0 {
		return nil, 0, nil
	}
	return items, 0, nil
}
