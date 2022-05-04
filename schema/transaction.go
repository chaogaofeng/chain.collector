package schema

import "time"

// Transaction defines the structure for transaction information.
type Transaction struct {
	ID        uint      `json:"-" gorm:"primarykey" `
	Height    int64     `json:"height" gorm:"not null"`
	TxHash    string    `json:"tx_hash" gorm:"not null,unique"`
	Code      uint32    `json:"code" gorm:"not null"`
	RawLog    string    `json:"raw_log" `
	Memo      string    `json:"memo"`
	Fees      string    `json:"fees" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	GasWanted int64     `json:"gas_wanted" gorm:"default:0"`
	GasUsed   int64     `json:"gas_used" gorm:"default:0"`
	Timestamp time.Time `json:"timestamp"`
}
