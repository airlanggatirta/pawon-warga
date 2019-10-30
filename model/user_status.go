package model

import "database/sql"

type UserStatus struct {
	ID          uint64         `gorm:"column:id"`
	Status      string         `gorm:"column:status"`
	Description sql.NullString `gorm:"column:description"`
}
