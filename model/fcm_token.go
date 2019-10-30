package model

type FcmToken struct {
	ID       uint64 `gorm:"column:id"`
	Platform string `gorm:"column:platform"`
	UserID   uint64 `gorm:"column:users_id"`
	Token    string `gorm:"column:token"`
	IsActive bool   `gorm:"column:is_active"`
	Created  int64  `gorm:"column:created"`
}

func (FcmToken) TableName() string {
	return "fcm_tokens"
}
