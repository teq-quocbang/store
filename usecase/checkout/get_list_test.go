package checkout

import (
	"context"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/store/delivery/http/auth"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/repository/checkout"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestGetListCart() {
	assertion := s.Assertions
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
		mockCheckout := checkout.NewMockRepository(s.T())
		mockCheckout.EXPECT().GetListCart(ctx, userPrinciple.User.ID).ReturnArguments = mock.Arguments{
			[]model.Cart{
				{
					AccountID: userPrinciple.User.ID,
					ProductID: testProductID,
					Qty:       int64(fake.Uint8()),
				},
			}, nil}
		u := s.useCase(mockCheckout)

		// Act
		reply, err := u.GetListCart(ctx)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}
}
