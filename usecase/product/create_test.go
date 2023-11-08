package product

import (
	"context"

	"bou.ke/monkey"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/repository/product"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
)

func (s *TestSuite) TestCreate() {
	assertion := s.Assertions
	testName := fake.Name()
	testProductType := fake.Car().Type
	testProducerID := uuid.New().String()

	userPrinciple := &token.JWTClaimCustom{
		SessionID: uuid.New(),
		User: token.UserInfo{
			ID:       uuid.New(),
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
		productModel := model.Product{
			Name:        testName,
			ProductType: testProductType,
			ProducerID:  uuid.MustParse(testProducerID),
			CreatedBy:   userPrinciple.User.ID,
		}
		mockProduct.EXPECT().Create(s.ctx, productModel).ReturnArguments = mock.Arguments{nil}
		u := s.useCase(mockProduct)
		req := &payload.CreateProductRequest{
			Name:        testName,
			ProductType: testProductType,
			ProducerID:  testProducerID,
		}

		// Act
		reply, err := u.Create(s.ctx, req)

		// Assert
		assertion.NoError(err)
		expected := &presenter.ProductResponseWrapper{
			Product: &model.Product{
				Name:        testName,
				ProductType: testProductType,
				ProducerID:  uuid.MustParse(testProducerID),
				CreatedBy:   userPrinciple.User.ID,
			},
		}
		assertion.Equal(expected, reply)
	}

	// bad case
	{ // missing name
		// Arrange
		req := &payload.CreateProductRequest{
			ProductType: testProductType,
			ProducerID:  testProducerID,
		}
		u := s.useCase(product.NewMockRepository(s.T()))

		// Act
		_, err := u.Create(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrProductInvalidParam("Key: 'CreateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
