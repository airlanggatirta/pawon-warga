package model

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type Topup struct {
	ID                 uint64            `gorm:"column:id"`
	TransactionID      uint64            `gorm:"column:transactions_id"`
	PaymentMethodID    uint64            `gorm:"column:payment_methods_id"`
	McBankName         sql.NullString    `gorm:"column:mc_bank_name"`
	McBankAccHolder    sql.NullString    `gorm:"column:mc_bank_acc_holder"`
	McAmount           sql.NullInt64     `gorm:"column:mc_amount"`
	McTransferDate     mysql.NullTime    `gorm:"column:mc_transfer_date"`
	McNotes            sql.NullString    `gorm:"column:mc_notes"`
	McImage            sql.NullString    `gorm:"column:mc_image"`
	McPOTRequestSent   bool              `gorm:"column:mc_pot_request_sent"`
	MutationScreenshot sql.NullString    `gorm:"column:screenshot_mutasi"`
	Created            int64             `gorm:"column:created"`
	Updated            int64             `gorm:"column:updated"`
	Expired            sql.NullInt64     `gorm:"column:expired"`
	Verified           sql.NullInt64     `gorm:"column:verified"`
	VaNumber           sql.NullString    `gorm:"column:va_number"`
	Fee                int64             `gorm:"column:fee"`
	PaymentMethod      PaymentMethod     `gorm:"foreignkey:PaymentMethodID"`
	WalletTransaction  WalletTransaction `gorm:"foreignkey:TransactionID"`
}

func (Topup) TableName() string {
	return "topups"
}

func (t Topup) ReferenceName() string {
	return strings.Title(t.TableName())
}
