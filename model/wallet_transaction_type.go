package model

import "database/sql"

type WalletTransactionType struct {
	ID          uint64         `gorm:"column:id"`
	Description sql.NullString `gorm:"column:description"`
}

func (WalletTransactionType) TableName() string {
	return "wallet_transaction_types"
}
