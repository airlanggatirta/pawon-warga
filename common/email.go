package common

import (
	"bytes"
	"html/template"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/airlanggatirta/pawon-warga/metric"
)

type IMailer interface {
	SendEmail(mailOption MailOption) error
}

type MailOption struct {
	Subject      string
	ToAddresses  []string
	CcAddresses  []string
	BccAddresses []string
	ContentHtml  string
	ContentText  string
}

type EmailTemplate struct {
	Header  *template.Template
	Content *template.Template
	Footer  *template.Template
}

const (
	// The character encoding for the email.
	CharSet = "UTF-8"
)

type SMTPMailerOption struct {
	Host       string
	Region     string
	Username   string
	Password   string
	Port       int
	SecureMode string
	Sender     string
	SenderName string
	ReplyTo    string
}

type SMTPMailer struct {
	Config     SMTPMailerOption
	sesSession *ses.SES
}

func setEnvVar(region string, accessKeyID string, secretKey string) {
	os.Setenv("AWS_REGION", region)
	os.Setenv("AWS_ACCESS_KEY_ID", accessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretKey)
}

func NewSMTPMailer(config SMTPMailerOption) (*SMTPMailer, error) {
	setEnvVar(config.Region, config.Username, config.Password)

	awsSession, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	sesSession := ses.New(awsSession)

	return &SMTPMailer{
		Config:     config,
		sesSession: sesSession,
	}, nil
}

func (sm *SMTPMailer) SendEmail(mailOption MailOption) error {
	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses:  aws.StringSlice(mailOption.CcAddresses),
			BccAddresses: aws.StringSlice(mailOption.BccAddresses),
			ToAddresses:  aws.StringSlice(mailOption.ToAddresses),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(mailOption.ContentHtml),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(mailOption.ContentText),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(mailOption.Subject),
			},
		},
		Source: aws.String(sm.Config.Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	// Capture start time for metrics
	startTime := time.Now()
	_, err := sm.sesSession.SendEmail(input)
	if err != nil {
		return err
	}
	endTime := time.Since(startTime)
	metric.ObserveSesDuration("send_email", endTime.Seconds())
	metric.CountSesPublished("email")
	return nil
}

func RenderEmail(template *template.Template, data interface{}) (string, error) {
	var buff bytes.Buffer
	err := template.Execute(&buff, data)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func LoadEmailTemplate(filePath string) *template.Template {
	template, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}

	return template
}
