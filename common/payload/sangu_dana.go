package payload

import (
	"strconv"
)

type Dana struct {
	PayReturnUrl string `json:"pay_return_url,omitempty" valid:"required"`
	ExpiredAt    string `json:"expired_at,omitempty" valid:"required"`
}

type DanaPayload struct {
	Vendor      string `json:"vendor"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	URLCallback string `json:"url_callback"`
	Desc        string `json:"desc"`
	Dana        Dana   `json:"dana"`
}

type DanaPayloadConfig struct {
	Currency       string
	Amount         float64
	URLCallback    string
	PayReturnUrl   string
	ExpiredAt      string
	Desc           string
	IdempotencyKey uint64
}

func NewDanaPayload(config DanaPayloadConfig) *DanaPayload {
	return &DanaPayload{
		Vendor:      "dana",
		Currency:    config.Currency,
		Amount:      strconv.FormatFloat(config.Amount, 'f', -1, 64),
		URLCallback: config.URLCallback,
		Desc:        config.Desc,
		Dana: Dana{
			PayReturnUrl:  config.PayReturnUrl,
			ExpiredAt:     config.ExpiredAt,
		},
	}
}
