package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository/account"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase func(*account.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(account *account.MockRepository) UseCase {
	return UseCase{
		Account: account,
		Config:  config.GetConfig(),
	}
}
