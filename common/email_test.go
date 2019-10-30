package common

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var randomSource rand.Source = rand.NewSource(time.Now().UnixNano())
var random *rand.Rand = rand.New(randomSource)

// TODO: make this unit test variables configurable

func GenerateActivationURL(userID uint64) string {
	activationKey := generateActivationKey()
	url := fmt.Sprintf("https://www.kitabisa.com/user/activation/%d/%s", userID, activationKey)
	return url
}

func generateActivationKey() string {
	now := time.Now().String()
	randomValue := random.Intn(100000000)
	activationKey := fmt.Sprintf("%s%d%s", "adsfsafsf", randomValue, now)

	return fmt.Sprintf("%x", sha1.Sum([]byte(activationKey)))
}

func TestSendEmail(t *testing.T) {
	mailerOption := SMTPMailerOption{
		Host:       "email-smtp.us-east-1.amazonaws.com",
		Region:     "us-east-1",
		Username:   "AKIAJC427LPHEMP5EBKA",
		Password:   "ynFALEZ/0k7gWymaj4tL3dOcmgE78NBQKMaDXwYx",
		Port:       587,
		SecureMode: "tls",
		Sender:     "support@kitabisa.com",
		SenderName: "Kitabisa.com",
		ReplyTo:    "support@kitabisa.com",
	}

	activationURL := GenerateActivationURL(125367)
	content := `
            <p>Halo <strong>Budi</strong>,</p>

            <p>Kami senang Anda bergabung di lingkaran kebaikan ini.</p>

            <p><strong><a href="` + activationURL + `">Aktifkan akun Anda</a></strong> untuk mulai menggalang dana atau berdonasi.</p>

            <p>Salam,</p>

            <p>Tim kitabisa.com</p>
			`

	mailer, err := NewSMTPMailer(mailerOption)
	assert.Nil(t, err, err)

	mailOption := MailOption{
		Subject:     "Test aja",
		ToAddresses: []string{"budi@kitabisa.com"},
		ContentHtml: content,
	}
	err = mailer.SendEmail(mailOption)
	assert.Nil(t, err, err)

}
