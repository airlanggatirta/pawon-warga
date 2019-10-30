package model

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Disbursement struct {
	ID                           uint64         `gorm:"column:id"`
	TransactionID                uint64         `gorm:"column:transactions_id"`
	Provider                     string         `gorm:"column:provider"`
	DestinationBankCode          sql.NullString `gorm:"column:dest_bank_code"`
	DestinationBankName          string         `gorm:"column:dest_bank_name"`
	DestinationBankBranch        sql.NullString `gorm:"column:dest_bank_branch"`
	DestinationBankAccountHolder string         `gorm:"column:dest_bank_acc_holder"`
	DestinationBankAccountNumber string         `gorm:"column:dest_bank_acc_number"`
	ActivationCode               string         `gorm:"column:activation_code"`
	VerificationMethod           string         `gorm:"column:verification_method"`
	Verified                     bool           `gorm:"column:verified"`
	OtpRequestCount              int            `gorm:"column:otp_request_count"`
	Created                      string         `gorm:"column:created"`
	Updated                      string         `gorm:"column:updated"`
	ProofOfTransfer              sql.NullString `gorm:"column:proof_of_transfer"`
	TransferAmount               int64          `gorm:"column:transfer_amount"`
	TransferFeeAmount            int64          `gorm:"column:transfer_fee_amount"`
	TotalDeducted                int64          `gorm:"column:total_deducted"`
	Status                       sql.NullString `gorm:"column:status"`
	AdminID                      sql.NullInt64  `gorm:"column:admin_id"`
	CreatedAt                    mysql.NullTime `gorm:"column:created_at"`
	UpdatedAt                    mysql.NullTime `gorm:"column:updated_at"`
}

// func (Disbursement) TableName() string {
// 	return "disbursements"
// }
