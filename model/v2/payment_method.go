package v2

type PaymentMethod struct {
	ID          string   `json:"id" mapstructure:"id"`
	Type        string   `json:"type" mapstructure:"type"`
	Category    string   `json:"category" mapstructure:"category"`
	Name        string   `json:"name" mapstructure:"name"`
	Desc        string   `json:"desc" mapstructure:"desc"`
	OffID       string   `json:"off_id" mapstructure:"off_id"`
	FeeVariable string   `json:"fee_variable" mapstructure:"fee_variable"`
	FeeFixed    string   `json:"fee_fixed" mapstructure:"fee_fixed"`
	Detail1     string   `json:"detail_1" mapstructure:"detail_1"`
	Detail2     string   `json:"detail_2" mapstructure:"detail_2"`
	Detail3     string   `json:"detail_3" mapstructure:"detail_3"`
	Detail4     string   `json:"detail_4" mapstructure:"detail_4"`
	Flag        []string `json:"flag" mapstructure:"flag"`
}

type Bank struct {
	Code       string `json:"code" mapstructure:"code"`
	Name       string `json:"name" mapstructure:"name"`
	Fee        int64  `json:"fee" mapstructure:"fee"`
	Status     int    `json:"status" mapstructure:"status"`
	StatusDesc string `json:"status_desc" mapstructure:"status_desc"`
}

type ResponseWithTotal struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type BankAccountInquiryReq struct {
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
}

type BankAccountInqruiryResp struct {
	AccountName string        `json:"account_name"`
	Status      InquiryStatus `json:"status"`
}

type InquiryStatus struct {
	Code int    `json:"code"`
	Desc string `json:"desc`
}
