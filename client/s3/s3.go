package s3

import (
	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/teq-quocbang/store/config"
)

var (
	sessionS3    *session.Session
	svcS3        *s3.S3
	downloaderS3 *s3manager.Downloader
	uploaderS3   *s3manager.Uploader
)

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSConfig.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSConfig.AccessKey,
			cfg.AWSConfig.SecretKey,
			"",
		),
	})
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	downloaderS3 = s3manager.NewDownloader(s)
	uploaderS3 = s3manager.NewUploader(s)

	svcS3 = s3.New(s)
}

func GetSession() *session.Session {
	return sessionS3
}

func GetService() *s3.S3 {
	return svcS3
}

func GetDownloader() *s3manager.Downloader {
	return downloaderS3
}

func GetUploader() *s3manager.Uploader {
	return uploaderS3
}
