package utils

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/opsdata-io/opsdata/pkg/config"
)

// GenerateDownloadLink generates a download link for a file in S3
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
