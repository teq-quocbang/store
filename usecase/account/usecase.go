package account

import (
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/repository/account"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Account account.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Account: repo.Account,
		SES:     ses,
		Config:  config.GetConfig(),
	}
}
