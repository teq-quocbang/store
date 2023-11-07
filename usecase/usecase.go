package usecase

import (
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase/account"
	"github.com/teq-quocbang/store/usecase/class"
	"github.com/teq-quocbang/store/usecase/course"
	"github.com/teq-quocbang/store/usecase/example"
	"github.com/teq-quocbang/store/usecase/grpc"
	"github.com/teq-quocbang/store/usecase/register"
	"github.com/teq-quocbang/store/usecase/semester"
	myS3 "github.com/teq-quocbang/store/util/s3"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Account  account.IUseCase
	Semester semester.IUseCase
	Class    class.IUseCase
	Course   course.IUseCase
	Register register.IUseCase
	Example  example.IUseCase
	GRPC     grpc.IUseCase

	SES mySES.ISES
	S3  myS3.IS3
}

func New(repo *repository.Repository, cache cache.ICache) *UseCase {
	var (
		ses = mySES.NewSES()
		s3  = myS3.NewS3()
	)

	return &UseCase{
		Account:  account.New(repo, ses),
		Class:    class.New(repo, ses),
		Course:   course.New(repo, ses),
		Semester: semester.New(repo, ses),
		Register: register.New(repo, ses, cache),
		Example:  example.New(repo, ses),
		GRPC:     grpc.New(repo),
		SES:      ses,
		S3:       s3,
	}
}
