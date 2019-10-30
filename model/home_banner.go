package model

type HomeBanner struct {
	ID           uint64 `gorm:"column:id"`
	Image        string `gorm:"column:image"`
	RedirectType string `gorm:"column:redirect_type"`
	RedirectID   string `gorm:"column:redirect_id"`
	IsActive     bool   `gorm:"column:is_active"`
	Priority     int64  `gorm:"column:priority"`
	Created      uint64 `gorm:"column:created"`
	Updated      uint64 `gorm:"column:updated"`
}

func (HomeBanner) TableName() string {
	return "apps_banners"
}
