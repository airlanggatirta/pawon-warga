package model

import "database/sql"

type CampaignCategory struct {
	ID             uint64         `gorm:"column:id"`
	NameID         string         `gorm:"column:name_id"`
	DescriptionID  sql.NullString `gorm:"column:description_id"`
	IsActive       bool           `gorm:"column:is_active"`
	Created        int64          `gorm:"column:created"`
	CmpnAccessRule string         `gorm:"column:cmpn_access_rule"`
	Icon           string         `gorm:"column:icon"`
	Slug           string         `gorm:"column:slug"`
	TotalDonation  uint64         `gorm:"column:total_donation"`
}

func (CampaignCategory) TableName() string {
	return "project_categories"
}
