package storage

import (
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/repository/producer"
	"github.com/teq-quocbang/store/repository/storage"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Producer producer.Repository
	Storage  storage.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Producer: repo.Producer,
		SES:      ses,
		Storage:  repo.Storage,
		Config:   config.GetConfig(),
	}
}
