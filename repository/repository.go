package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/store/repository/account"
)

type Repository struct {
	Account account.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Account: account.NewAccountPG(getClient),
	}
}
