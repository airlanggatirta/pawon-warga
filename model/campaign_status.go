package model

type CampaignStatus struct {
	ID          uint64 `gorm:"column:id"`
	Status      string `gorm:"column:status"`
	Description string `gorm:"column:description"`
	IsEditable  bool   `gorm:"column:is_editable"`
}

func (CampaignStatus) TableName() string {
	return "project_statuses"
}
