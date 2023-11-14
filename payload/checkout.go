package payload

import "github.com/go-playground/validator/v10"

type AddToCartRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Qty       int64  `json:"qty" validate:"required"`
}

func (a AddToCartRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}

type GetCartRequest struct {
	ProductID string `json:"product_id" validate:"required"`
}

func (c GetCartRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
