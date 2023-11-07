package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	myS3 "github.com/teq-quocbang/store/client/s3"
)

type S3 struct {
	session    *session.Session
	svc        *s3.S3
	downloader *s3manager.Downloader
	uploader   *s3manager.Uploader
}

func NewS3() IS3 {
	return S3{
		session:    myS3.GetSession(),
		svc:        myS3.GetService(),
		downloader: myS3.GetDownloader(),
		uploader:   myS3.GetUploader(),
	}
}
