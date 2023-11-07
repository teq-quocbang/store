package account

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	SignUp(context.Context, *payload.SignUpRequest) (*presenter.AccountResponseWrapper, error)
	Login(context.Context, *payload.LoginRequest) (*presenter.AccountLoginResponseWrapper, error)
}
