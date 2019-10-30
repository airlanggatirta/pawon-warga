package common

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
)

var WhatsappBaseURL string

func WhatsappValidate(WaNumber string) (err error) {
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

	requestURL := fmt.Sprintf("%s/%s", WhatsappBaseURL, WaNumber)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			return ErrWhatsappNumberNotRegistered
		case http.StatusInternalServerError:
			return errors.New(http.StatusText(http.StatusInternalServerError))
		}
	}

	return nil
}
