package product

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository/product"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func(*product.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(product *product.MockRepository) UseCase {
	return UseCase{
		Product: product,
		Config:  config.GetConfig(),
	}
}
