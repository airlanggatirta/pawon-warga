package service

import "github.com/airlanggatirta/pawon-warga/model"

type IUserService interface {
	GetOneUserByEmailWithoutFilter(email string) (model.User, error)
	GetOneUserByWhatsappWithoutFilter(waNumber string) (user model.User, err error)
}
