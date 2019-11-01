package service

import (
	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/repository"
)

type UserServiceBuilder struct {
	UserRepository           repository.IUserRepository
	Option                   ServiceOption
}

func NewUserServiceBuilder(option ServiceOption) *UserServiceBuilder {
	return &UserServiceBuilder{Option: option}
}

func (usb *UserServiceBuilder) SetUserRepository(userRepo repository.IUserRepository) *UserServiceBuilder {
	usb.UserRepository = userRepo
	return usb
}