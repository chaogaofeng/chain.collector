package schema

import (
	"gorm.io/gorm"
	"time"
)

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

func QueryTxByHash(db *gorm.DB, hash string) (*Transaction, error) {
	var item Transaction
	if res := db.Model(&Transaction{}).Where("tx_hash = ?", hash).Find(&item); res.Error != nil {
		return nil, res.Error
	} else if res.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}
