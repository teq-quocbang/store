package producer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/store/config"
	"github.com/teq-quocbang/store/repository/producer"
)

type TestSuite struct {
	suite.Suite

	ctx context.Context

	useCase func(*producer.MockRepository) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(producer *producer.MockRepository) UseCase {
	return UseCase{
		Producer: producer,
		Config:   config.GetConfig(),
	}
}
