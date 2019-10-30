package model

import "database/sql"

type WalletTransaction struct {
	ID              uint64         `gorm:"column:id"`
	UserID          uint64         `gorm:"column:users_id"`
	ModelID         sql.NullInt64  `gorm:"column:model_id"`
	Model           sql.NullString `gorm:"column:model"`
	TypesID         uint64         `gorm:"column:types_id"`
	Amount          int64          `gorm:"column:amount"`
	Flow            string         `gorm:"column:flow"`
	Status          string         `gorm:"column:status"`
	Source          string         `gorm:"column:source"`
	Description     string         `gorm:"column:description"`
	Created         int64          `gorm:"column:created"`
	Updated         int64          `gorm:"column:updated"`
	Name            string         `grom:"column:name"`
	UserDescription string         `grom:"column:user_description"`
	Disbursement    Disbursement   `gorm:"foreignkey:ID;association_foreignkey:TransactionID"`
}

func (WalletTransaction) TableName() string {
	return "wallet_transactions"
}
