package product

import (
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/repository/product"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Product product.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Product: repo.Product,
		SES:     ses,
		Config:  config.GetConfig(),
	}
}
