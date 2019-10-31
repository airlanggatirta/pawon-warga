package driver

import (
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
)

type CloudSearchOption struct {
	Key      string
	Secret   string
	Endpoint string
	Region   string
}

func NewCloudSearch(session *session.Session) (csd *cloudsearchdomain.CloudSearchDomain) {
	csd = cloudsearchdomain.New(session)
	return
}

func AuthCloudSearch(option CloudSearchOption) (sess *session.Session, err error) {
	var token string
	credential := credentials.NewStaticCredentials(option.Key, option.Secret, token)

	config := &aws.Config{
		Endpoint:    &option.Endpoint,
		Region:      &option.Region,
		Credentials: credential,
	}

	sess, err = session.NewSession(config)
	if err != nil {
		return sess, err
	}

	return sess, nil
}
