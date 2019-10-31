package service

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/model"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/asaskevich/govalidator.v4"
)

// userService defines struct for V3 user service
type userService struct {
	userService IUserService
	UserServiceOption
}

// userServiceOption defines struct for V3 user service options
type UserServiceOption struct {
	ServiceOption
}

// RegisterTemp will do user temporary register
func (us *userService) RegisterTemp(user model.User, uuidToken string) (string, int64, error) {
	// TODO: input validation

	cache := us.Cache.Get()
	defer cache.Close()

	cacheTTL := 0
	isSendActivationEmail := false

	if user.Email.String != "" {
		userExists, err := us.userService.IsUserExists(user.Email.String)
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return "", 0, err
		}

		if userExists {
			return "", 0, common.ErrUserAlreadyExists
		}

		user.TempID = us.hashTempID(user.Email.String)
		err = us.getRegistrationTempTTL(user.TempID)
		if err != nil {
			return "", 0, err
		}

		user.ActivationKey = us.generateActivationKey()
		cacheTTL = 172800
		isSendActivationEmail = true
	} else if user.WaPhoneNumber.String != "" {
		userExists, err := us.userService.IsUserExists(user.WaPhoneNumber.String)
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return "", 0, err
		}

		if userExists {
			return "", 0, common.ErrUserAlreadyExists
		}

		user.TempID = us.hashTempID(user.WaPhoneNumber.String)
		err = us.getRegistrationTempTTL(user.TempID)

		if err != nil && err != common.ErrPendingRegistration {
			return "", 0, err
		}

		if err == common.ErrPendingRegistration {
			// TODO: complete this condition
			return "", 0, err
		}

		user.ActivationKey = us.generateActivationKey()
		cacheTTL = 900
	} else {
		return "", 0, common.ErrEmptyUserID
	}

	data, err := json.Marshal(user)
	if err != nil {
		return "", 0, err
	}

	_, err = cache.Do("SETEX", fmt.Sprintf("%s:%s", "registration", user.TempID), cacheTTL, data)
	if err != nil {
		return "", 0, err
	}

	if isSendActivationEmail {
		err := us.userService.SendActivationEmail(user, user.Email.String)
		if err != nil {
			return "", 0, err
		}
		return "", 0, nil
	}

	ttl := int64(0)
	return uuidToken, ttl, nil
}

func (us *userService) generateActivationKey() string {
	now := time.Now().String()
	randomValue := random.Intn(100000000)
	activationKey := fmt.Sprintf("%s%d%s", us.AppConfig.Users.ActivationKeySalt, randomValue, now)

	return fmt.Sprintf("%x", sha1.Sum([]byte(activationKey)))
}

func (us *userService) hashTempID(data string) string {
	hash := []byte("735kh325" + data + "c23s4h53k")
	return fmt.Sprintf("%x", sha1.Sum(hash))
}

func (us *userService) getRegistrationTempTTL(data string) error {
	cache := us.Cache.Get()
	defer cache.Close()

	ttl, err := redis.Int64(cache.Do("TTL", fmt.Sprintf("%s:%s", "registration", data)))
	if err != nil {
		return err
	}

	if ttl > 0 {
		return common.ErrPendingRegistration
	}

	return nil
}

func (us *userService) IsUserExists(username string) (exists bool, err error) {
	if govalidator.IsEmail(username) {
		return us.checkUser(us.GetUserByEmail, username)
	} else if govalidator.Matches(username, "^(0|62|\\+)\\d+$") { // TODO: support i18n for country code number
		phoneNumber := common.ReplacePhoneNumberMask(username)
		return us.checkUser(us.GetUserByWANumber, phoneNumber)
	} else {
		err = common.ErrInvalidUserID
		return
	}
}

func (us *userService) SendActivationEmail(user model.User, emailAddress string) error {
	activationURL, err := us.GenerateActivationURL(user.TempID, user.ActivationKey)
	if err != nil {
		us.Logger.HandleErrorWithTrace(err)
		return err
	}

	contentData := EmailContentData{
		FullName:       user.FullName,
		ActivationLink: activationURL,
	}

	header, err := common.RenderEmail(us.RegistrationEmailTemplate.Header, nil)
	if err != nil {
		us.Logger.HandleErrorWithTrace(err)
		return err
	}

	content, err := common.RenderEmail(us.RegistrationEmailTemplate.Content, contentData)
	if err != nil {
		us.Logger.HandleErrorWithTrace(err)
		return err
	}

	footer, err := common.RenderEmail(us.RegistrationEmailTemplate.Footer, nil)
	if err != nil {
		us.Logger.HandleErrorWithTrace(err)
		return err
	}

	fullEmailContent := fmt.Sprintf("%s%s%s", header, content, footer)

	mailOption := common.MailOption{
		Subject:     "Aktivasi Akun Anda",
		ToAddresses: []string{emailAddress},
		ContentHtml: fullEmailContent,
	}

	err = us.smtpMailer.SendEmail(mailOption)
	if err != nil {
		us.Logger.HandleErrorWithTrace(err)
		return err
	}

	return nil
}
