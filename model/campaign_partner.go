package model

import "database/sql"

type CampaignPartner struct {
	ID               uint64         `gorm:"column:id"`
	UserID           uint64         `gorm:"column:users_id"`
	Name             string         `gorm:"column:name"`
	Description      string         `gorm:"column:description"`
	URL              sql.NullString `gorm:"column:url"`
	Image            sql.NullString `gorm:"column:image"`
	Slug             string         `gorm:"column:slug"`
	IsActive         bool           `gorm:"column:is_active"`
	Created          int64          `gorm:"column:created"`
	Updated          int64          `gorm:"column:updated"`
	DisplayOrder     int64          `gorm:"column:display_order"`
	ProjectsOrdering string         `gorm:"column:projects_ordering"`
	CtaText          sql.NullString `gorm:"column:cta_text"`
	CtaURL           sql.NullString `gorm:"column:cta_url"`
	PartnerType      sql.NullString `gorm:"column:partner_type"`
	BannerImage      sql.NullString `gorm:"column:banner_image"`
	User             User
}

func (CampaignPartner) TableName() string {
	return "partners"
}
