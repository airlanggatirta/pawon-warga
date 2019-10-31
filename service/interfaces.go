package service

import (
	"github.com/airlanggatirta/pawon-warga/model"
)

type IUserService interface {
	GetAllUsers() ([]model.User, error)
	GetOneUser(id uint64) (model.User, error)
	GetOneUserBySecondaryID(id string) (model.User, error)
	GetOneUserByCredential(userID string, password string) (model.User, error)
	GetOneUserByActivationKey(userID string, activationKey string) (model.User, error)
	GetOneUserByEmailAndStatusActive(emailAddress string) (model.User, error)
	GetOneUserByWhatsapp(waNumber string) (model.User, error)
	GetOneNonLoginUserByEmailAddress(email string) (*model.User, error)
	GetOneNonLoginUserByWANumber(waNumber string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint64) error
	DeleteActivationKey(userID string, activationKey string) error
	HashPassword(password string) (string, error)
	IsVerified(user model.User) bool
	IsPasswordMatch(password, hash string) bool
	UpdateUserAPIToken(user model.User, token string) error
	SearchUser(keyword string, pageToken string, campaignerOnly bool) (users []model.User, nextToken string, err error)
	GetOneUserByEmailWithoutFilter(email string) (user model.User, err error)
	GetOneUserByWhatsappWithoutFilter(waNumber string) (user model.User, err error)
	IsUserExists(username string) (exists bool, err error)
	SendActivationEmail(user model.User, emailAddress string) error
}
