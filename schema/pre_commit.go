package schema

import "time"

// PreCommit defines the schema for precommit state information.
type PreCommit struct {
	ID               uint      `json:"-" gorm:"primarykey" `
	Height           int64     `json:"height" grom:"not null"`
	Round            int32     `json:"round" grom:"not null"`
	ValidatorAddress string    `json:"validator_address" grom:"not null"`
	VotingPower      int64     `json:"voting_power" grom:"not null"`
	ProposerPriority int64     `json:"proposer_priority" grom:"not null"`
	Timestamp        time.Time `json:"timestamp"`
}
