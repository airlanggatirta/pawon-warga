package model

import (
	"database/sql"
	"strings"
)

type Lovelist struct {
	ID         	uint64			`gorm:"column:id"`
	UserID     	uint64 			`gorm:"column:users_id"`
	CampaignID 	uint64 			`gorm:"column:campaign_id"`
	Created    	int64  			`gorm:"column:created"`
	Updated    	int64			`gorm:"column:updated"`
	DeletedAt	sql.NullInt64	`gorm:"column:deleted"`
	Campaign   	Campaign
}

func (Lovelist) TableName() string {
	return "lovelists"
}

func (t Lovelist) ReferenceName() string {
	return strings.Title(t.TableName())
}
