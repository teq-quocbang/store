package payload

import "github.com/go-playground/validator/v10"

type CreateProducerRequest struct {
	Name    string `json:"name" validate:"required"`
	Country string `json:"country" validate:"required"`
}

func (p *CreateProducerRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(p)
}
