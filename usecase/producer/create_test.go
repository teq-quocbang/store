package producer

import (
	"context"

	"bou.ke/monkey"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/producer"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions
	testName := fake.Name()
	testCountry := fake.Country()

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID:       uuid.New(),
			Email:    "test@teqnological.asia",
			Username: "test_username",
		},
	}
	monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})
	defer monkey.UnpatchAll()

	// good case
	{
		// Arrange
		mockProducer := producer.NewMockRepository(s.T())
		producerModel := &model.Producer{
			Name:      testName,
			Country:   testCountry,
			CreatedBy: userPrinciple.User.ID,
			UpdatedBy: userPrinciple.User.ID,
		}
		mockProducer.EXPECT().Create(s.ctx, producerModel).ReturnArguments = mock.Arguments{nil}
		u := s.useCase(mockProducer)
		req := &payload.CreateProducerRequest{
			Name:    testName,
			Country: testCountry,
		}

		// Act
		reply, err := u.Create(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{ // missing name
		// Arrange
		req := &payload.CreateProducerRequest{
			Country: testCountry,
		}
		u := s.useCase(producer.NewMockRepository(s.T()))

		// Act
		_, err := u.Create(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrProducerInvalidParam("Key: 'CreateProducerRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
