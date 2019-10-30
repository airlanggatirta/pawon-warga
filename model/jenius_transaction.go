package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type JeniusTransaction struct {
	Id                      uint64  `gorm:"column:id"`
	Model                   string  `gorm:"column:model"`
	ModelID                 uint64  `gorm:"column:model_id"`
	TransactionID           string  `gorm:"column:transaction_id"`
	TransactionCreatedTime  int64   `gorm:"column:transaction_created_time"`
	TransactionPaidTime     int64   `gorm:"column:transaction_paid_time"`
	TransactionNotifiedTime int64   `gorm:"column:transaction_notified_time"`
	Currency                string  `gorm:"column:currency"`
	Amount                  float64 `gorm:"column:amount"`
	Description             string  `gorm:"column:description"`
	VendorReferenceNumber   string  `gorm:"column:vendor_reference_no"`
	StatusCode              int     `gorm:"column:status_code"`
	Cashtag                 string  `gorm:"column:cashtag"`
	PromoCode               string  `gorm:"column:promo_code"`
	Created                 int64   `gorm:"column:created"`
}

func (JeniusTransaction) TableName() string {
	return "jenius_transactions"
}

func (u *JeniusTransaction) BeforeCreate(scope *gorm.Scope) (err error) {
	u.Created = time.Now().Unix()
	return
}
