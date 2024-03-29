package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/jinzhu/gorm"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/model"
	"github.com/airlanggatirta/pawon-warga/service"
	"github.com/airlanggatirta/pawon-warga/version"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code       int
	Err        error
	ThirdParty bool
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type HttpHandler struct {
	Logger *common.APILogger
	H      func(w http.ResponseWriter, r *http.Request) error
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(w, r)
	if err != nil {

		switch e := err.(type) {
		case Error:
			if ee, ok := e.(StatusError); ok {

				if gorm.IsRecordNotFoundError(ee.Err) {
					emptyResp := NewAPIResponse(nil, nil)

					resp, err := json.Marshal(emptyResp)
					if err != nil {
						WriteErrorResponse(w, ee)
						return
					}

					w.WriteHeader(http.StatusOK)
					w.Write(resp)

					return
				}

				WriteErrorResponse(w, ee)
			}
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

type APIErrorResponse struct {
	APICode int        `json:"api_code"`
	Errors  []APIError `json:"errors"`
}
type APIError struct {
	Code    *int    `json:"code,omitempty"`
	Details Details `json:"details"`
}

type Details struct {
	Id string `json:"id"`
	En string `json:"en"`
}

type APIResponse struct {
	APICode int         `json:"api_code"`
	Next    *string     `json:"next,omitempty"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
}

type Meta struct {
	Version string `json:"version"`
	Status  string `json:"api_status"`
	APIEnv  string `json:"api_env"`
}

type Handler struct {
	Service    *service.Service
	JWTSignKey string
	Logger     *common.APILogger
}

func NewHandler(service *service.Service, signKey string, logger *common.APILogger) *Handler {
	return &Handler{
		Service:    service,
		JWTSignKey: signKey,
		Logger:     logger,
	}
}

func NewAPIResponse(data interface{}, nextPage *string) APIResponse {
	var response APIResponse
	voData := reflect.ValueOf(data)
	arrayData := []interface{}{}
	if voData.Kind() != reflect.Slice {
		if voData.IsValid() {
			arrayData = []interface{}{data}
		}
		response.Data = arrayData
	} else {
		if voData.Len() != 0 {
			response.Data = data
		} else {
			response.Data = arrayData
		}
	}
	response.APICode = common.GetAPICode(nil)
	response.Meta.Version = version.Version
	response.Meta.APIEnv = version.Environment
	response.Meta.Status = "unstable"
	response.Next = nextPage
	return response
}

func NewAPIErrorResponse(err error, apiErr ...APIError) APIErrorResponse {
	apiErrorResponse := APIErrorResponse{}
	apiErrorResponse.APICode = common.GetAPICode(err)
	apiErrorResponse.Errors = apiErr
	return apiErrorResponse
}

func GetRequestBody(r *http.Request) (reqBody []byte, err error) {
	reqBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, common.ErrInvalidRequestEntity
	}

	return
}

func ParseRequestBody(reqBody []byte, target interface{}) (err error) {
	err = json.Unmarshal(reqBody, target)

	if err != nil {
		return common.ErrMissingRequestEntity
	}

	return
}

func WriteErrorResponse(w http.ResponseWriter, err StatusError) {
	w.WriteHeader(err.Status())

	apiError := APIError{}

	if err.ThirdParty {
		apiError.Code = &err.Code
		apiError.Details = Details{
			Id: err.Err.Error(),
			En: err.Err.Error(),
		}
	} else if _, ok := errorMapEnglish[err.Err]; !ok {
		apiError.Details = Details{
			Id: "Ups ada kesalahan, silakan coba dalam beberapa saat lagi.",
			En: err.Err.Error(),
			// En: "Oops something went wrong, try again later."
		}
	} else {
		apiError.Details = Details{
			Id: errorMapBahasa[err.Err],
			En: errorMapEnglish[err.Err],
		}
	}

	apiErrResponse := NewAPIErrorResponse(err.Err, apiError)
	errResp, _ := json.Marshal(apiErrResponse)
	w.Write(errResp)
}

func (h *Handler) ParseUserFromContext(ctx context.Context) (user model.User, err error) {
	var userID uint64
	fmt.Println(userID)
	claims, ok := ctx.Value("UserInfo").(*model.NewKitabisaClaims)
	if !ok {
		userInfo, _ := ctx.Value("UserInfo").(model.User)
		userID = userInfo.ID
	} else {
		userID = claims.UserID
	}

	return
}
