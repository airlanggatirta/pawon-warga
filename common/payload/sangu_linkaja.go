package payload

import (
	"strconv"
)

type LinkAja struct {
	SuccessUrl  string `json:"success_url,omitempty" valid:"required"`
	FailUrl     string `json:"fail_url,omitempty" valid:"required"`
	PhoneNumber string `json:"phone_number,omitempty" valid:"required"`
}

type LinkAjaPayload struct {
	Vendor      string  `json:"vendor"`
	Currency    string  `json:"currency"`
	Amount      string  `json:"amount"`
	URLCallback string  `json:"url_callback"`
	Desc        string  `json:"desc"`
	LinkAja     LinkAja `json:"linkaja"`
}

type LinkAjaPayloadConfig struct {
	Currency       string
	Amount         float64
	URLCallback    string
	SuccessUrl     string
	FailUrl        string
	PhoneNumber    string
	Desc           string
	IdempotencyKey uint64
}

func NewLinkAjaPayload(config LinkAjaPayloadConfig) *LinkAjaPayload {
	return &LinkAjaPayload{
		Vendor:      "linkaja",
		Currency:    config.Currency,
		Amount:      strconv.FormatFloat(config.Amount, 'f', -1, 64),
		URLCallback: config.URLCallback,
		Desc:        config.Desc,
		LinkAja: LinkAja{
			SuccessUrl:  config.SuccessUrl,
			FailUrl:     config.FailUrl,
			PhoneNumber: config.PhoneNumber,
		},
	}
}
