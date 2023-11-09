package producer

import (
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/repository/producer"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Producer producer.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Producer: repo.Producer,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
