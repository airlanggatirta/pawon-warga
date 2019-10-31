package service

import (
	"math/rand"
	"time"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/config"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var randomSource rand.Source = rand.NewSource(time.Now().UnixNano())
var random *rand.Rand = rand.New(randomSource)

// ServiceV3Option defines V3 options
type ServiceOption struct {
	DB        *gorm.DB          `validate:"required"`
	Cache     *redis.Pool       `validate:"required"`
	Logger    *common.APILogger `validate:"required"`
	AppConfig *config.Config    `validate:"required"`
}

// ServiceV3 defines struct for V3 service
type Service struct {
	User IUserService
}
