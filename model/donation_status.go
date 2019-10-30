package model

type DonationStatus struct {
	ID          uint64 `gorm:"column:id"`
	Status      string `gorm:"column:status"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
