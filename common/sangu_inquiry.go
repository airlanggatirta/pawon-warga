package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/spf13/viper"
)

type InquiryResponse struct {
	TransactionID     string   `json:"id_transaction"`
	Vendor            string   `json:"vendor"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
	PaidAt            string   `json:"paid_at"`
	NotifiedAt        string   `json:"notified_at"`
	Currency          string   `json:"currency"`
	Amount            string   `json:"amount"`
	Desc              string   `json:"desc"`
	VendorReferenceNo string   `json:"vendor_reference_no"`
	Midtrans          Midtrans `json:"midtrans"`
	Status            Status   `json:"status"`
}

type Midtrans struct {
	FirstName       string            `json:"first_name"`
	LastName        string            `json:"last_name"`
	Email           string            `json:"email"`
	Phone           string            `json:"phone"`
	MidtransActions []MidtransActions `json:"actions"`
}

type MidtransActions struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

type Status struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type SanguRespError struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

func GetInquiry(token string, vendor_trx_id string) (inquiry InquiryResponse, err error) {
	backoffInterval := 2 * time.Millisecond
	maximumJitterInterval := 5 * time.Millisecond
	backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)

	retrier := heimdall.NewRetrier(backoff)

	timeout := 1000 * time.Millisecond

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(3),
	)

	bearer := "Bearer " + token

	sanguBaseURL := viper.GetString("sangu.sangu_base_url")

	requestURL := fmt.Sprintf("%s/pay/%s", sanguBaseURL, vendor_trx_id)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return inquiry, err
	}

	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		return inquiry, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return inquiry, err
	}

	err = json.Unmarshal(body, &inquiry)
	if err != nil {
		return inquiry, err
	}

	defer resp.Body.Close()

	return
}
