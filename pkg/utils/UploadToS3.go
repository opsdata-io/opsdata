package utils

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/opsdata-io/opsdata/pkg/config"
)

// UploadToS3 uploads a file to S3
func UploadToS3(file []byte, filename string) error {
	svc := s3.New(s3Session)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.CFG.S3Bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})

	return err
}
