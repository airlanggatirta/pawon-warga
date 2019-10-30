package model

type Privilege struct {
	ID          uint64 `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
