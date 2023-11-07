package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/teq-quocbang/store/config"
)

func (s S3) Delete(key string) (*s3.DeleteObjectOutput, error) {
	cfg := config.GetConfig()

	result, err := s.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(cfg.S3Config.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
