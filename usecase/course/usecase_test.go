package course

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository/course"
	"github.com/teq-quocbang/store/repository/semester"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase func(*course.MockRepository, *semester.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(course *course.MockRepository, semester *semester.MockRepository) UseCase {
	return UseCase{
		Course:   course,
		Semester: semester,
		Config:   config.GetConfig(),
	}
}
