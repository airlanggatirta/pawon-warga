package model

import (
	"github.com/go-sql-driver/mysql"
)

// CampaignMedicalDocument defines model for project_medical_documents table
type CampaignMedicalDocument struct {
	ID         uint64         `gorm:"column:id" db:"id"`
	ProjectsID uint64         `gorm:"column:projects_id" db:"projects_id"`
	Path       string         `gorm:"column:path" db:"path"`
	StatusID   uint16         `gorm:"column:status_id" db:"status_id"`
	UserID     uint64         `gorm:"column:user_id" db:"user_id"`
	Reason     string         `gorm:"column:reason" db:"reason"`
	CreatedAt  mysql.NullTime `gorm:"column:created_at" db:"created_at"`
	UpdatedAt  mysql.NullTime `gorm:"column:updated_at" db:"updated_at"`
	DeletedAt  mysql.NullTime `gorm:"column:deleted_at" db:"deleted_at"`
}

// TableName returns table name for project_medical_documents
func (CampaignMedicalDocument) TableName() string {
	return "project_medical_documents"
}
