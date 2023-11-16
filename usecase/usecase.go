package usecase

import (
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/usecase/account"
	"github.com/teq-quocbang/store/usecase/checkout"
	"github.com/teq-quocbang/store/usecase/example"
	"github.com/teq-quocbang/store/usecase/grpc"
	"github.com/teq-quocbang/store/usecase/producer"
	"github.com/teq-quocbang/store/usecase/product"
	"github.com/teq-quocbang/store/usecase/statistics"
	"github.com/teq-quocbang/store/usecase/storage"
	myS3 "github.com/teq-quocbang/store/util/s3"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Account    account.IUseCase
	Example    example.IUseCase
	GRPC       grpc.IUseCase
	Product    product.IUseCase
	Producer   producer.IUseCase
	Storage    storage.IUseCase
	Checkout   checkout.IUseCase
	Statistics statistics.IUseCase

	SES mySES.ISES
	S3  myS3.IS3
}

func New(repo *repository.Repository, cache cache.ICache) *UseCase {
	var (
		ses = mySES.NewSES()
		s3  = myS3.NewS3()
	)

	return &UseCase{
		Account:    account.New(repo, ses),
		Example:    example.New(repo, ses),
		GRPC:       grpc.New(repo),
		Product:    product.New(repo, ses),
		Producer:   producer.New(repo, ses),
		Storage:    storage.New(repo, ses),
		Checkout:   checkout.New(repo, ses),
		Statistics: statistics.New(repo, ses),
		SES:        ses,
		S3:         s3,
	}
}
