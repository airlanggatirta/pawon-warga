package model

type UniqueCodeReservation struct {
	ID               uint64 `gorm:"column:id"`
	Amount           int64  `gorm:"column:amount"`
	Code             int64  `gorm:"column:code"`
	Bank             string `gorm:"column:bank"`
	ReferenceModel   string `gorm:"column:model"`
	ReferenceModelID uint64 `gorm:"column:model_id"`
	Created          int64  `gorm:"column:created"`
	Expired          int64  `gorm:"column:expired"`
}
