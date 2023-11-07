package example_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	exampleRepo "github.com/teq-quocbang/store/repository/example"
	exampleUC "github.com/teq-quocbang/store/usecase/example"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase *exampleUC.UseCase

	mockExampleRepo *exampleRepo.MockRepository
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.mockExampleRepo = &exampleRepo.MockRepository{}

	suite.useCase = &exampleUC.UseCase{
		ExampleRepo: suite.mockExampleRepo,
	}
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
