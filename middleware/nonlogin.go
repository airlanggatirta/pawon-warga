package middleware

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/handler"
	"github.com/airlanggatirta/pawon-warga/model"
	"gopkg.in/asaskevich/govalidator.v4"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	TypeEmail = "Email"
	TypePhone = "Phone"
)

type NonLoginRequest struct {
	Name         string `json:"name"`
	EmailOrPhone string `json:"email_or_phone"`
}

func getIsEmailOrPhone(emailOrPhone string) string {
	ok := govalidator.IsEmail(emailOrPhone)
	if !ok {
		validatorPhone := regexp.MustCompile(`^\+?\d{9,14}$`)
		validatePhone := validatorPhone.MatchString(emailOrPhone)
		if validatePhone {
			return TypePhone
		}
	}
	return TypeEmail
}

func (m *Middleware) NonLoginMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	body, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		statusError := handler.StatusError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
		handler.WriteErrorResponse(w, statusError)
		return
	}

	var request NonLoginRequest

	err = json.Unmarshal(body, &request)
	if err != nil {
		statusError := handler.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
		handler.WriteErrorResponse(w, statusError)
		return
	}

	if request.Name == "" {
		statusError := handler.StatusError{
			Code: http.StatusBadRequest,
			Err:  common.ErrCreateDonationNonLoginEmptyName,
		}
		handler.WriteErrorResponse(w, statusError)
		return
	}

	if request.EmailOrPhone == "" {
		statusError := handler.StatusError{
			Code: http.StatusBadRequest,
			Err:  common.ErrCreateDonationNonLoginEmptyEmailWaNumber,
		}
		handler.WriteErrorResponse(w, statusError)
		return
	}

	if request.EmailOrPhone != "" {
		ok := govalidator.IsEmail(request.EmailOrPhone)
		if !ok {
			validatorPhone := regexp.MustCompile(`^\+?\d{9,14}$`)
			validatePhone := validatorPhone.MatchString(request.EmailOrPhone)
			if !validatePhone {
				statusError := handler.StatusError{
					Code: http.StatusBadRequest,
					Err:  common.ErrCreateDonationNonLoginInvalidEmailWaNumber,
				}
				handler.WriteErrorResponse(w, statusError)
				return
			}
		}
	}

	var user model.User
	var userNotExist bool
	var userID uint64

	ctx := r.Context()
	getTypeContact := getIsEmailOrPhone(request.EmailOrPhone)

	switch getTypeContact {
	case TypeEmail:
		user, err = m.UserService.GetOneUserByEmailWithoutFilter(request.EmailOrPhone)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				userNotExist = true
				user.Email = sql.NullString{String: request.EmailOrPhone, Valid: true}
			}
		}
	case TypePhone:
		user, err = m.UserService.GetOneUserByWhatsappWithoutFilter(common.ReplacePhoneNumberMask(request.EmailOrPhone))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				userNotExist = true
				user.WaPhoneNumber = sql.NullString{String: common.ReplacePhoneNumberMask(request.EmailOrPhone), Valid: true}
			}
		}
	}

	if userNotExist {
		user.FullName = request.Name
		userDummy, err := m.UserService.RegisterDummy(&user)
		if err != nil {
			statusError := handler.StatusError{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
			handler.WriteErrorResponse(w, statusError)
			return
		}
		userID = userDummy.ID
	} else {
		userID = user.ID
	}

	UserInfo := &model.NewKitabisaClaims{
		UserID: userID,
	}

	ctx = context.WithValue(ctx, "UserInfo", UserInfo)

	next(w, r.WithContext(ctx))
}
