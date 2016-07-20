package sms_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	env "github.com/segmentio/go-env"
	"github.com/stretchr/testify/assert"

	"github.com/tj/go-sms"
)

var number = env.MustGet("PHONE")

func TestSMS_Send(t *testing.T) {
	service := sns.New(session.New(aws.NewConfig()))

	sms := &sms.SMS{
		Service:  service,
		Type:     sms.Transactional,
		MaxPrice: 0.50,
	}

	err := sms.Send("Hello from SMS.Send()", number)
	assert.NoError(t, err, "error sending")
}

func TestSend(t *testing.T) {
	err := sms.Send("Hello from Send()", number)
	assert.NoError(t, err, "error sending")
}
