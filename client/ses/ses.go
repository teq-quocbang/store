package ses

import (
	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/teq-quocbang/store/config"
)

var (
	sessionSES *session.Session
	svcSES     *ses.SES
)

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	sessionSES, err = session.NewSession(&aws.Config{
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

	svcSES = ses.New(sessionSES)
}

func GetSession() *session.Session {
	return sessionSES
}

func GetService() *ses.SES {
	return svcSES
}
