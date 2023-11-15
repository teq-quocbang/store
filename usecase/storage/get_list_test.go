package storage

import (
	"context"
	"fmt"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/delivery/http/auth"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository/storage"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestGetListByLocat() {
	assertion := s.Assertions
	testLocat := fmt.Sprintf("%s%d", "A", fake.IntRange(100, 1000))
	testQty := fake.Int8()
	testProductID := uuid.New()

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID:       uuid.New(),
			Email:    "test@teqnological.asia",
			Username: "test_username",
		},
	}
	ctx := context.WithValue(s.ctx, auth.UserPrincipleKey, userPrinciple)

	// good case
	{
		// Arrange
		mockStorage := storage.NewMockRepository(s.T())
		mockStorage.EXPECT().GetListStorageByLocat(ctx, testLocat).ReturnArguments = mock.Arguments{
			[]model.Storage{
				{
					Locat:        testLocat,
					ProductID:    testProductID,
					InventoryQty: int64(testQty),
				},
			}, nil}
		req := &payload.GetStorageByLocatRequest{
			Locat: testLocat,
		}
		u := s.useCase(mockStorage)

		// Act
		reply, err := u.GetList(ctx, req)

		// Assert
		assertion.NoError(err)
		expected := &presenter.ListStorageResponseWrapper{
			Storage: []model.Storage{
				{
					Locat:        testLocat,
					ProductID:    testProductID,
					InventoryQty: int64(testQty),
				},
			},
		}
		assertion.Equal(expected, reply)
	}

	// bad case
	{
		// Arrange
		u := s.useCase(storage.NewMockRepository(s.T()))

		// Act
		reply, err := u.GetList(ctx, &payload.GetStorageByLocatRequest{})

		// Assert
		assertion.NoError(err) // a special return with return null according business logic
		assertion.Equal(&presenter.ListStorageResponseWrapper{}, reply)
	}
}
