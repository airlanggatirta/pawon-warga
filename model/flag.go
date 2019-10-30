package model

import "database/sql"

type Flag struct {
	ID             uint64         `gorm:"column:id"`
	ReferenceModel string         `gorm:"column:model"`
	Name           string         `gorm:"column:name"`
	Description    sql.NullString `gorm:"column:description"`
	IsActive       bool           `gorm:"column:is_active"`
	Created        int64          `gorm:"column:created"`
	Updated        int64          `gorm:"column:updated"`
}

func (Flag) TableName() string {
	return "flags"
}
