package account

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	CreateAccount(context.Context, *model.Account) (ID uint, err error)
	GetAccountByID(ctx context.Context, studentID uint) (*model.Account, error)
	GetAccountByConstraint(context.Context, *model.Account) (*model.Account, error)
	CreateVerifyAccount(context.Context, *model.AccountVerify) error
	GetVerifyAccountByID(ctx context.Context, studentID uint) (*model.AccountVerify, error)
	GetList(context.Context) ([]model.Account, error)
}
