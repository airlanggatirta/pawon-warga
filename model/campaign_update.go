package model

import "database/sql"

type CampaignUpdate struct {
	ID                       uint64                   `gorm:"column:id"`
	ProjectsID               uint64                   `gorm:"column:projects_id"`
	Title                    string                   `gorm:"column:title"`
	Content                  string                   `gorm:"column:content"`
	Image                    sql.NullString           `gorm:"column:image"`
	Source                   sql.NullString           `gorm:"column:source"`
	Created                  int64                    `gorm:"column:created"`
	Published                sql.NullInt64            `gorm:"column:published"`
	GoodImages               sql.NullString           `gorm:"column:good_images"`
	BadImages                sql.NullString           `gorm:"column:bad_images"`
	IsAdminPublished         sql.NullBool             `gorm:"column:is_admin_published"`
	ContentImage             string                   `gorm:"-"`
	DisbursementUpdateDetail DisbursementUpdateDetail `gorm:"-"`
	Campaign                 Campaign                 `gorm:"foreignkey:ProjectsID"`
}

type CampaignUpdateFilter struct {
	ProjectsID uint64 // query campaign updates by campaign_id
}

func (CampaignUpdate) TableName() string {
	return "project_updates"
}

type DisbursementUpdateDetail struct {
	DestinationBankName          string `gorm:"column:dest_bank_name"`
	DestinationAccountHolderName string `gorm:"column:dest_bank_acc_holder"`
	DestinationAccountNumber     string `gorm:"column:dest_bank_acc_number"`
}
