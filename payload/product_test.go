package payload

import (
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductRequest(t *testing.T) {
	assertion := assert.New(t)

	// good case
	{
		// Arrange
		req := CreateProductRequest{
			Name:        fake.Name(),
			ProductType: fake.Car().Type,
			ProducerID:  uuid.New().String(),
		}

		// Act
		err := req.Validate()

		// Assert
		assertion.NoError(err)
	}
	// missing name
	{
		// Arrange
		req := CreateProductRequest{
			ProductType: fake.Car().Type,
			ProducerID:  uuid.New().String(),
		}

		// Act
		err := req.Validate()

		// Assert
		assertion.Error(err)
		assertion.Equal("Key: 'CreateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag", err.Error())
	}
	// missing product type
	{
		// Arrange
		req := CreateProductRequest{
			Name:       fake.Name(),
			ProducerID: uuid.New().String(),
		}

		// Act
		err := req.Validate()

		// Assert
		assertion.Error(err)
		assertion.Equal("Key: 'CreateProductRequest.ProductType' Error:Field validation for 'ProductType' failed on the 'required' tag", err.Error())
	}
	// missing producer id
	{
		// Arrange
		req := CreateProductRequest{
			Name:        fake.Name(),
			ProductType: fake.Car().Type,
		}

		// Act
		err := req.Validate()

		// Assert
		assertion.Error(err)
		assertion.Equal("Key: 'CreateProductRequest.ProducerID' Error:Field validation for 'ProducerID' failed on the 'required' tag", err.Error())
	}
}
