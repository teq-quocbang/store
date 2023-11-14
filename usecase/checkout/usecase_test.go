package checkout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository/checkout"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func(*checkout.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(checkout *checkout.MockRepository) UseCase {
	return UseCase{
		Checkout: checkout,
		Config:   config.GetConfig(),
	}
}
