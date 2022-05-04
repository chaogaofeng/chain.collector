package schema

import sdk "github.com/cosmos/cosmos-sdk/types"

type Account struct {
	ID                  uint     `json:"-" gorm:"primarykey" `
	Address             string   `json:"address"`
	AccountNumber       uint64   `json:"account_number"`
	Total               sdk.Coin `json:"total"`
	Balance             sdk.Coin `json:"balance"`
	Delegation          sdk.Coin `json:"delegation"`
	UnbondingDelegation sdk.Coin `json:"unbonding_delegation"`
	Rewards             sdk.Coin `json:"rewards"`
	UpdateAt            int64    `json:"update_at"`
	CreateAt            int64    `json:"create_at"`
}
