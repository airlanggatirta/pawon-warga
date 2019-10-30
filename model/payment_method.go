package model

import "database/sql"

type PaymentMethod struct {
	ID             uint64         `gorm:"column:id"`
	Type           sql.NullString `gorm:"column:type"`
	Name           sql.NullString `gorm:"column:name"`
	Category       string         `gorm:"column:category"`
	Shortname      sql.NullString `gorm:"column:shortname"`
	Detail1        sql.NullString `gorm:"column:detail_1"`
	Detail2        sql.NullString `gorm:"column:detail_2"`
	Detail3        sql.NullString `gorm:"column:detail_3"`
	Detail4        sql.NullString `gorm:"column:detail_4"`
	IsTopup        bool           `gorm:"column:is_topup"`
	VariableFee    float64        `gorm:"column:variable_fee"`
	FixedFee       int64          `gorm:"column:fixed_fee"`
	Image          sql.NullString `gorm:"column:image"`
	DisplayOrder   int64          `gorm:"column:display_order"`
	IsHidden       bool           `gorm:"column:is_hidden"`
	IsActive       bool           `gorm:"column:is_active"`
	FlipIdentifier sql.NullString `gorm:"column:flip_identifier"`
	Description    sql.NullString `gorm:"column:description"`
	DesktopOnly    bool           `gorm:"column:desktop_only"`
}
