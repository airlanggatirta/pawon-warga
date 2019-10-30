package model

type UserDevice struct {
	ID          uint64 `gorm:"column:id"`
	Platform    string `gorm:"column:platform"`
	UserID      uint64 `gorm:"column:users_id"`
	DeviceToken string `gorm:"column:device_token"`
	FcmToken    string `gorm:"column:fcm_token"`
	ClevertapID string `gorm:"column:clevertap_id"`
	Created     int64  `gorm:"column:created"`
	Updated     int64  `gorm:"column:updated"`
}

func (UserDevice) TableName() string {
	return "user_devices"
}
