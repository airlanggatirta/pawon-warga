package model

type Province struct {
	ID   uint64 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
