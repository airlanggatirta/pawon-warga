package model

import "database/sql"

type VaPairedTransaction struct {
	ID                       uint64        `gorm:"column:id"`
	Model                    string        `gorm:"column:model"`
	ModelID                  uint64        `gorm:"column:model_id"`
	TransactionReferenceCode string        `gorm:"column:transaction_reference_code"`
	BankCode                 string        `gorm:"column:bank_code"`
	VaNumber                 string        `gorm:"column:va_number"`
	Created                  sql.NullInt64 `gorm:"column:created"`
}
