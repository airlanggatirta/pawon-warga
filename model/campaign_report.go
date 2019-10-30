package model

type CampaignReport struct {
	ID           uint64   `gorm:"column:id"`
	CampaignID   uint64   `gorm:"column:projects_id"`
	UserID       uint64   `gorm:"column:users_id"`
	Content      string   `gorm:"column:contents"`
	SeverityFlag int64    `gorm:"column:severity_flag"`
	Status       bool     `gorm:"column:status"`
	Created      int64    `gorm:"column:created"`
	Updated      int64    `gorm:"column:updated"`
	Campaign     Campaign //`gorm:"foreignkey:ProjectsID"`
	User         User     //`gorm:"foreignkey:UserID"`
}

func (CampaignReport) TableName() string {
	return "project_reports"
}
