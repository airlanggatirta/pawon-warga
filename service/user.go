package service

import "github.com/airlanggatirta/pawon-warga/model"

// userService defines struct for V3 user service
type userService struct {
	userService IUserService
	UserServiceOption
}

// userServiceOption defines struct for V3 user service options
type UserServiceOption struct {
	ServiceOption
}

func (us *userService) GetOneUserByEmailWithoutFilter(email string) (model.User, error) {
	return us.repository.GetOneUserByEmailWithoutFilter(email)
}

func (us *userService) GetOneUserByWhatsappWithoutFilter(waNumber string) (user model.User, err error) {
	return us.repository.GetOneUserByWhatsappWithoutFilter(waNumber)
}
