package register

import (
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository"
	"github.com/teq-quocbang/store/repository/account"
	"github.com/teq-quocbang/store/repository/class"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/register"
	"github.com/teq-quocbang/store/repository/semester"
	mySES "github.com/teq-quocbang/store/util/ses"
)

type UseCase struct {
	Account  account.Repository
	Semester semester.Repository
	Class    class.Repository
	Course   course.Repository
	Register register.Repository

	SES mySES.ISES

	Cache cache.ICache

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES, cache cache.ICache) IUseCase {
	return &UseCase{
		Account:  repo.Account,
		Semester: repo.Semester,
		Class:    repo.Class,
		Course:   repo.Course,
		Register: repo.Register,
		Cache:    cache,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
