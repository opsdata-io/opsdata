package utils

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/opsdata-io/opsdata/backend/pkg/config"
)

// DownloadFromS3 downloads a file from S3 and returns the file as a byte array
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
	if _, err := buf.ReadFrom(result.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
