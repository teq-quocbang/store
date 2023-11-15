package payload

import "github.com/go-playground/validator/v10"

type UpsertStorageRequest struct {
	Locat     string `json:"locat" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Qty       int64  `json:"qty" validate:"required"`
}

func (s UpsertStorageRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(s)
}

type GetStorageByLocatRequest struct {
	Locat string `json:"locat,omitempty" query:"locat" validate:"required"`
}

func (s GetStorageByLocatRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(s)
}
