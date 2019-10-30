package model

import "database/sql"

type CampaignMarketingCampaign struct {
	ID            uint64         `gorm:"column:id"`
	ProjectsID    uint64         `gorm:"column:projects_id"`
	Code          string         `gorm:"column:code"`
	Description   string         `gorm:"column:description"`
	FeePercentage float64        `gorm:"column:fee_percentage"`
	URL           sql.NullString `gorm:"column:url"`
	UtmCampaign   sql.NullString `gorm:"column:utm_campaign"`
	UtmMedium     sql.NullString `gorm:"column:utm_medium"`
	UtmSource     sql.NullString `gorm:"column:utm_source"`
	RefererID     sql.NullInt64  `gorm:"column:referer_id"`
	Created       int64          `gorm:"column:created"`
	Updated       int64          `gorm:"column:updated"`
	Campaign      Campaign
}

func (CampaignMarketingCampaign) TableName() string {
	return "project_marketing_campaign"
}
