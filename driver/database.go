package driver

import (
	"fmt"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBOption struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	ConnectionSetting    ConnectionSetting
}

type ConnectionSetting struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewMysqlDatabase(option DBOption) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
	if err != nil {
		return nil, err
	}

	sqlDB := db.DB()
	sqlDB.SetConnMaxLifetime(option.ConnectionSetting.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(option.ConnectionSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(option.ConnectionSetting.MaxOpenConns)

	return db, nil
}
