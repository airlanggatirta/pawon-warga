package payload

import "strconv"

type Jenius struct {
	Cashtag   string `json:"cashtag"`
	PromoCode string `json:"promo_code"`
}

type JeniusPayload struct {
	Vendor      string `json:"vendor"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	URLCallback string `json:"url_callback"`
	Desc        string `json:"desc"`
	Jenius      Jenius `json:"jenius"`
}

type JeniusPayloadConfig struct {
	Cashtag     string
	Currency    string
	Amount      float64
	CallbackURL string
	Desc        string
	PromoCode   string
}

func NewJeniusPayload(config JeniusPayloadConfig) *JeniusPayload {
	return &JeniusPayload{
		Vendor:      "btpn",
		Currency:    config.Currency,
		Amount:      strconv.FormatFloat(config.Amount, 'f', -1, 64),
		URLCallback: config.CallbackURL,
		Desc:        config.Desc,
		Jenius: Jenius{
			Cashtag:   config.Cashtag,
			PromoCode: config.PromoCode,
		},
	}
}
