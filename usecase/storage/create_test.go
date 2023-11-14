package storage

import (
	"context"
	"fmt"

	"bou.ke/monkey"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/storage"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestUpsertStorage() {
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
	monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})
	defer monkey.UnpatchAll()

	// good case
	{
		// Arrange
		mockStorage := storage.NewMockRepository(s.T())
		mockStorage.EXPECT().UpsertStorage(s.ctx, &model.Storage{
			Locat:        testLocat,
			ProductID:    testProductID,
			InventoryQty: int64(testQty),
			CreatedBy:    userPrinciple.User.ID,
			UpdatedBy:    userPrinciple.User.ID,
		}).ReturnArguments = mock.Arguments{nil}
		req := &payload.UpsertStorageRequest{
			Locat:     testLocat,
			ProductID: testProductID.String(),
			Qty:       int64(testQty),
		}
		u := s.useCase(mockStorage)

		// Act
		reply, err := u.UpsertStorage(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}
}
