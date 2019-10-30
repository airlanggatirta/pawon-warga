package model

import "database/sql"

type FlagAssignment struct {
	ID       uint64        `gorm:"column:id"`
	FlagsID  uint64        `gorm:"column:flags_id"`
	ModelID  uint64        `gorm:"column:model_id"`
	AdminID  sql.NullInt64 `gorm:"column:admin_id"`
	Status   bool          `gorm:"column:status"`
	IsRecent bool          `gorm:"column:is_recent"`
	Created  int64         `gorm:"column:created"`
}

func (FlagAssignment) TableName() string {
	return "flag_assignments"
}
