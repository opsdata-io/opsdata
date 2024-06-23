package utils

import (
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/opsdata-io/opsdata/pkg/config"
)

var s3Session *session.Session

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

func UploadToS3(file []byte, filename string) error {
	svc := s3.New(s3Session)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.CFG.S3Bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})

	return err
}

func DownloadFromS3(filename string) ([]byte, error) {
	svc := s3.New(s3Session)

	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.CFG.S3Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	return buf.Bytes(), nil
}

func GenerateDownloadLink(filename string) (string, error) {
	svc := s3.New(s3Session)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.CFG.S3Bucket),
		Key:    aws.String(filename),
	})

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}
