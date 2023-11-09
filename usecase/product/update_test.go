package product

import (
	"context"

	"bou.ke/monkey"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/product"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestUpdate() {
	assertion := s.Assertions
	testName := fake.Name()
	testProductType := fake.Car().Type
	testProducerID := uuid.New().String()

	productID := uuid.New()
	userID := uuid.New()
	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID:       userID,
			Email:    "test@teqnological.asia",
			Username: "test_username",
		},
	}
	user := monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return userPrinciple
	})
	defer monkey.Unpatch(user)

	// good case
	{
		// Arrange
		mockProduct := product.NewMockRepository(s.T())
		productModel := &model.Product{
			ID:          productID,
			Name:        testName,
			ProductType: testProductType,
			ProducerID:  uuid.MustParse(testProducerID),
			UpdatedBy:   userID,
		}
		mockProduct.EXPECT().Update(s.ctx, productModel).ReturnArguments = mock.Arguments{nil}
		mockProduct.EXPECT().GetByID(s.ctx, productID).ReturnArguments = mock.Arguments{model.Product{
			ID:          productID,
			Name:        fake.Name(),
			ProductType: fake.Car().Type,
			ProducerID:  uuid.MustParse(testProducerID),
		}, nil}
		u := s.useCase(mockProduct)

		// Act
		reply, err := u.Update(s.ctx, &payload.UpdateProductRequest{
			ID:          productID.String(),
			Name:        testName,
			ProductType: testProductType,
			ProducerID:  testProducerID,
		})

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{
		u := s.useCase(product.NewMockRepository(s.T()))

		// Act
		_, err := u.Update(s.ctx, &payload.UpdateProductRequest{
			Name:        testName,
			ProductType: testProductType,
			ProducerID:  testProducerID,
		})

		// Assert
		assertion.Error(err)
	}
}
