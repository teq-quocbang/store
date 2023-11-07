package semester

import (
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/repository/semester"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Semester semester.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Semester: repo.Semester,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
