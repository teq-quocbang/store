package product

import (
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"

	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/repository/product"
)

func (s *TestSuite) TestExport() {
	assertion := s.Assertions

	// good case
	{
		// Arrange
		mockProduct := product.NewMockRepository(s.T())
		mockProduct.EXPECT().GetList(s.ctx).ReturnArguments = mock.Arguments{
			[]model.Product{
				{
					Name:        fake.Name(),
					ProductType: fake.Car().Type,
					ProducerID:  uuid.New(),
				},
				{
					Name:        fake.Name(),
					ProductType: fake.Car().Type,
					ProducerID:  uuid.New(),
				},
				{
					Name:        fake.Name(),
					ProductType: fake.Car().Type,
					ProducerID:  uuid.New(),
				},
				{
					Name:        fake.Name(),
					ProductType: fake.Car().Type,
					ProducerID:  uuid.New(),
				},
				{
					Name:        fake.Name(),
					ProductType: fake.Car().Type,
					ProducerID:  uuid.New(),
				},
			}, nil}
		req := &payload.ExportProductRequest{
			FileExtension: "csv",
		}
		u := s.useCase(mockProduct)

		// Act
		reply, err := u.Export(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply.Meta)
	}
}
