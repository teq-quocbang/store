package payload

import (
	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}

func (s *SignUpRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
