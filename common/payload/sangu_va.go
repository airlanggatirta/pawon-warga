package payload

import (
	"strconv"
)

type VA struct {
	Type 		int		`json:"type"`
	VANumber 	string	`json:"va_number"`
	ExpiredAt	string	`json:"expired_at"`

}

type VAPayload struct {
	Vendor		string 	`json:"vendor"`
	Currency	string	`json:"currency"`
	Amount		string	`json:"amount"`
	URLCallback	string	`json:"url_callback"`
	Desc		string	`json:"desc"`
	VA			VA		`json:"va"`
}

type VAPayloadConfig struct {
	Vendor			string
	Currency 		string
	Amount			float64
	URLCallback		string
	Desc			string
	Type			int
	VANumber		string
	ExpiredAt		string
	IdempotencyKey	uint64
}

func NewVAPayload(config VAPayloadConfig) *VAPayload {
	return &VAPayload{
		Vendor: config.Vendor,
		Currency: config.Currency,
		Amount: strconv.FormatFloat(config.Amount, 'f', -1, 64),
		URLCallback: config.URLCallback,
		Desc: config.Desc,
		VA: VA{
			Type: config.Type,
			VANumber: config.VANumber,
			ExpiredAt: config.ExpiredAt,
		},
	}
}