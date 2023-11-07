package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type IS3 interface {
	Upload(file io.Reader, dir RootDirectory) (*s3manager.UploadOutput, string, error)
	Delete(key string) (*s3.DeleteObjectOutput, error)
}
