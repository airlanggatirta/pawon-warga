package common

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/airlanggatirta/pawon-warga/metric"
)

type SNSOptions struct {
	Host     string
	Region   string
	Username string
	Password string
}

type SNSEvent struct {
	Config     SNSOptions
	snsSession *sns.SNS
}

func setSNSEnvVar(region string, accessKeyID string, secretKey string) {
	os.Setenv("AWS_SNS_REGION", region)
	os.Setenv("AWS_SNS_ACCESS_KEY_ID", accessKeyID)
	os.Setenv("AWS_SNS_SECRET_ACCESS_KEY", secretKey)
}

func NewEvent(config SNSOptions) (*SNSEvent, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.Username, config.Password, ""),
	})
	if err != nil {
		return nil, err
	}

	snsSession := sns.New(awsSession)

	return &SNSEvent{
		Config:     config,
		snsSession: snsSession,
	}, nil
}

func (e *SNSEvent) Publish(eventName string, payload map[string]interface{}, topicArn string) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	eventPayload := &sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Subject:  aws.String(eventName),
		Message:  aws.String(string(jsonPayload)),
	}
	// Capture start time for metrics
	startTime := time.Now()
	_, err = e.snsSession.Publish(eventPayload)
	if err != nil {
		return err
	}
	endTime := time.Since(startTime)
	metric.ObserveSnsDuration("publish", endTime.Seconds())
	metric.CountSnsPublished(eventName)
	return nil
}
