package model

type Location struct {
	ID         uint64 `gorm:"column:id"`
	Province   Province
	ProvinceID uint64 `gorm:"column:provinces_id"`
	Prefix     string `gorm:"column:prefix"`
	City       string `gorm:"column:city"`
}
