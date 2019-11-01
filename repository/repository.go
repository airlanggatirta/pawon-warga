package repository

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/config"
)

type RepositoryOption struct {
	DB        *gorm.DB                             `validate:"required"`
	Cache     *redis.Pool                          `validate:"required"`
	Logger    *common.APILogger                    `validate:"required"`
	AppConfig *config.Config                       `validate:"required"`
	CSD       *cloudsearchdomain.CloudSearchDomain `validate:"required"`
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}
