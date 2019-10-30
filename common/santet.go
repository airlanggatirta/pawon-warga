package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
)

var (
	// SantetBaseURL defines santet base url
	SantetBaseURL string
	// SantetUsername defines santet username (basic auth)
	SantetUsername string
	// SantetPassword defines santet password (basic auth)
	SantetPassword string
	// SantetBasicAuthToggle defines santet basic auth toggle
	SantetBasicAuthToggle bool
)

const (
	Clevertap = "clevertap"
	Sms       = "sms"
)

type Android struct {
	BackgroundImage string `json:"background_image"`
	DefaultSound    string `json:"default_sound"`
	DeepLink        string `json:"deep_link"`
	LargeIcon       string `json:"large_icon"`
	Key             string `json:"key"`
	WzrkCid         string `json:"wzrk_cid"`
}

type Ios struct {
	DeepLink   string `json:"deep_link"`
	SoundFile  string `json:"sound_file"`
	Category   string `json:"category"`
	BadgeCount string `json:"badge_count"`
	Key        string `json:"key"`
}

type Windows struct {
	DeepLink string `json:"deep_link"`
	Key      string `json:"key"`
}

type PlatformSpecific struct {
	Windows Windows `json:"windows"`
	Ios     Ios     `json:"ios"`
	Android Android `json:"android"`
}

type Content struct {
	Title            string           `json:"title"`
	Body             string           `json:"body"`
	PlatformSpecific PlatformSpecific `json:"platform_specific"`
}

type Body struct {
	RespectFrequencyCaps bool    `json:"respect_frequency_caps"`
	Content              Content `json:"content"`
}

type ClevertapPayload struct {
	Destination string `json:"destination"`
	Type        string `json:"type"`
	Body        Body   `json:"body"`
}

type Phone struct {
	RegionCode string `json:"region_code"`
}
type CustomParam struct {
	Phone Phone `json:"phone"`
}

type SmsPayload struct {
	Destination string      `json:"destination"`
	Body        string      `json:"body"`
	Type        string      `json:"type"`
	CustomParam CustomParam `json:"custom_param"`
}

func Santet(typeNotification string, destination string, title string, body string, deeplink string) (err error) {
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

	var payloadMarshal []byte

	switch typeNotification {
	case Clevertap:
		clevertapPayload := ClevertapPayload{}
		clevertapPayload.Destination = destination
		clevertapPayload.Body.Content.Title = title
		clevertapPayload.Body.Content.Body = body
		clevertapPayload.Body.Content.PlatformSpecific.Windows.DeepLink = deeplink
		clevertapPayload.Body.Content.PlatformSpecific.Ios.DeepLink = deeplink
		clevertapPayload.Body.Content.PlatformSpecific.Android.DeepLink = deeplink
		clevertapPayload.Body.Content.PlatformSpecific.Android.WzrkCid = "com.kitabisa.android"
		clevertapPayload.Type = Clevertap

		payloadMarshal, err = json.Marshal(clevertapPayload)
	case Sms:
		smsPayload := SmsPayload{}
		smsPayload.Destination = destination
		smsPayload.Body = body
		smsPayload.Type = Sms
		smsPayload.CustomParam.Phone.RegionCode = "ID"

		payloadMarshal, err = json.Marshal(smsPayload)
	}

	if err != nil {
		return err
	}

	requestURL := fmt.Sprintf("%s/message/notify", SantetBaseURL)
	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(payloadMarshal))
	if err != nil {
		return err
	}

	if SantetBasicAuthToggle {
		req.SetBasicAuth(SantetUsername, SantetPassword)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return
}
