package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository/storage"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func(*storage.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(storage *storage.MockRepository) UseCase {
	return UseCase{
		Storage: storage,
		Config:  config.GetConfig(),
	}
}
