package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type DonationReminder struct {
	ID           uint64 	`gorm:"column:id"`
	UserID       uint64 	`gorm:"column:users_id"`
	Frequency    string 	`gorm:"column:frequency"`
	Day          string 	`gorm:"column:day"`
	Time         string 	`gorm:"column:time"`
	TimeLocation string 	`gorm:"column:time_location"`
	IsUpdated    bool   	`gorm:"column:is_updated"`
	IsDeleted    bool   	`gorm:"column:is_deleted"`
	CreatedAt    int64  	`gorm:"column:created_at"`
	UpdatedAt    int64 		`gorm:"column:updated_at"`
	DeletedAt    *time.Time	`gorm:"column:deleted_at"`
}

func (DonationReminder) TableName() string {
	return "donation_reminders"
}

func (d *DonationReminder) BeforeCreate(scope *gorm.Scope) (err error) {
	d.CreatedAt = time.Now().Unix()
	return
}

func (d *DonationReminder) BeforeUpdate(scope *gorm.Scope) (err error) {
	d.UpdatedAt = time.Now().Unix()
	return
}
