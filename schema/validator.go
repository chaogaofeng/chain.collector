package schema

import (
	"gorm.io/gorm"
	"time"
)

// Validator defines the structure for validator information.
type Validator struct {
	ID                      uint      `json:"-" gorm:"primarykey" `
	Moniker                 string    `json:"moniker"`
	AccountAddress          string    `json:"account_address" gorm:"not null, unique"`
	OperatorAddress         string    `json:"operator_address" gorm:"not null, unique"`
	ConsensusAddress        string    `json:"consensus_address" gorm:"not null, unique"`
	Jailed                  bool      `json:"jailed"`
	Status                  string    `json:"status"`
	Tokens                  string    `json:"tokens"`
	VotingPower             int64     `json:"voting_power"`
	DelegatorShares         string    `json:"delegator_shares"`
	BondHeight              int64     `json:"bond_height" gorm:"default:0"`
	BondIntraTxCounter      int64     `json:"bond_intra_tx_counter" gorm:"default:0"`
	UnbondingHeight         int64     `json:"unbonding_height" gorm:"default:0"`
	UnbondingTime           string    `json:"unbonding_time"`
	CommissionRate          string    `json:"commission_rate"`
	CommissionMaxRate       string    `json:"commission_max_rate"`
	CommissionMaxChangeRate string    `json:"commission_max_change_rate"`
	CommissionUpdateTime    string    `json:"commission_update_time"`
	Timestamp               time.Time `json:"timestamp"`
}

// QueryValidatorMoniker returns validator's moniker.
func QueryValidator(db *gorm.DB, valAddr string) *Validator {
	var validator Validator
	if res := db.Model(&Validator{}).Where("consensus_address = ?", valAddr).Find(&validator); res.Error != nil {
		return nil
	} else if res.RowsAffected == 0 {
		return nil
	}
	return &validator
}
