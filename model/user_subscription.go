package model

type UserSubscription struct {
	ID      uint64 `gorm:"column:id"`
	UserID  uint64 `gorm:"column:users_id"`
	Model   string `gorm:"column:model"`
	ModelID int64  `gorm:"column:model_id"`
	User    User
}
