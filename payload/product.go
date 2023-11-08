package payload

import "github.com/go-playground/validator/v10"

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	ProductType string `json:"product_type" validate:"required"`
	ProducerID  string `json:"producer_id" validate:"required"`
}

func (p CreateProductRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}
