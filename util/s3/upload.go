package s3

import (
	"bytes"
	"io"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/satori/go.uuid"

	"github.com/teq-quocbang/store/config"
)

type RootDirectory string

const (
	RootDirectoryFiles  RootDirectory = "/files"
	RootDirectoryAvatar RootDirectory = "/avatars"
)

func (s S3) Upload(file io.Reader, dir RootDirectory) (*s3manager.UploadOutput, string, error) {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, "", err
	}

	var (
		contentType = http.DetectContentType(buf.Bytes())
		cfg         = config.GetConfig()
		fileName    = uuid.NewV5(uuid.NewV4(), cfg.S3Config.KeyUUID).String() + teq.FileTypeString(contentType)
	)

	pathFileName := string(dir) + fileName
	if cfg.S3Config.DefaultDir != "" {
		pathFileName = cfg.S3Config.DefaultDir + pathFileName
	}

	result, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(cfg.S3Config.BucketName),
		Key:         aws.String(pathFileName),
		Body:        buf,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return nil, "", err
	}

	return result, pathFileName, nil
}
