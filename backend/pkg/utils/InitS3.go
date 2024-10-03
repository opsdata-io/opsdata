package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/opsdata-io/opsdata/backend/pkg/config"
)

var s3Session *session.Session

// InitS3 initializes the S3 session
func InitS3() error {
	var err error
	s3Session, err = session.NewSession(&aws.Config{
		Region:           aws.String(config.CFG.S3Region),
		Endpoint:         aws.String(config.CFG.S3Endpoint),
		Credentials:      credentials.NewStaticCredentials(config.CFG.S3AccessKey, config.CFG.S3SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	return err
}
