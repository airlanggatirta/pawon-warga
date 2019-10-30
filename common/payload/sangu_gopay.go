package payload	

import (	
	"strconv"	
)	

type Midtrans struct {	
	FirstName			string	`json:"first_name"`	
	LastName			string	`json:"last_name"`	
	Email				string	`json:"email"`	
	Phone				string	`json:"phone"`	
	RedirectCallback	string	`json:"redirect_callback"`	
}	


type GopayPayload struct {	
	Vendor      string `json:"vendor"`	
	Currency    string `json:"currency"`	
	Amount      string `json:"amount"`	
	URLCallback string `json:"url_callback"`	
	Desc        string `json:"desc"`	
	Midtrans	Midtrans `json:"midtrans"`	
}	

type GopayPayloadConfig struct {	
	FirstName			string	
	LastName			string	
	Email				string	
	Phone				string	
	Currency			string	
	Amount				float64	
	URLCallback			string	
	RedirectCallback	string	
	Desc				string	
	IdempotencyKey		uint64	
}	

func NewGopayPayload(config GopayPayloadConfig) *GopayPayload {	
	return &GopayPayload{	
		Vendor: "gopay",	
		Currency: config.Currency,	
		Amount: strconv.FormatFloat(config.Amount, 'f', -1, 64),	      
		URLCallback: config.URLCallback,	
		Desc: config.Desc,	
		Midtrans: Midtrans{	
			FirstName: config.FirstName,	
			LastName: config.LastName,	
			Email: config.Email,	
			Phone: config.Phone,	
			RedirectCallback: config.RedirectCallback,	
		},	
	}
}	