package product

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/repository/product"
	"github.com/teq-quocbang/store/util/myerror"
)

func (s *TestSuite) TestDelete() {
	assertion := s.Assertions
	testProductID := uuid.New()

	// good case
	{
		// Arrange
		mockProduct := product.NewMockRepository(s.T())
		mockProduct.EXPECT().Delete(s.ctx, testProductID).ReturnArguments = mock.Arguments{nil}
		mockProduct.EXPECT().GetByID(s.ctx, testProductID).ReturnArguments = mock.Arguments{model.Product{}, nil}
		u := s.useCase(mockProduct)

		// Act
		err := u.Delete(s.ctx, testProductID.String())

		// Assert
		assertion.NoError(err)
	}
	// bad case
	{
		// Arrange
		u := s.useCase(product.NewMockRepository(s.T()))

		// Act
		err := u.Delete(s.ctx, "")

		// Assert
		assertion.Error(err)
		assertion.Equal(myerror.ErrProductInvalidParam("missing id"), err)
	}
}
