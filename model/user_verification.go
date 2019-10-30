package model

import "database/sql"

type UserVerification struct {
	ID                uint64         `gorm:"column:id"`
	UserID            uint64         `gorm:"column:users_id"`
	IDScan            sql.NullString `gorm:"column:id_scan"`
	IDNumber          sql.NullString `gorm:"column:id_number"`
	BankCode          sql.NullString `gorm:"column:bank_code"`
	BankName          sql.NullString `gorm:"column:bank_name"`
	BankAccountHolder sql.NullString `gorm:"column:bank_account_holder"`
	BankAccountNumber sql.NullString `gorm:"column:bank_account_number"`
	IsVerified        sql.NullBool   `gorm:"column:is_verified"`
	IsActive          sql.NullBool   `gorm:"column:is_active"`
	Created           sql.NullString `gorm:"column:created"`
	Updated           sql.NullString `gorm:"column:updated"`
	User              User
}
