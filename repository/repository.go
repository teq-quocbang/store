package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/store/repository/account"
	"github.com/teq-quocbang/store/repository/checkout"
	"github.com/teq-quocbang/store/repository/example"
	"github.com/teq-quocbang/store/repository/producer"
	"github.com/teq-quocbang/store/repository/product"
	"github.com/teq-quocbang/store/repository/storage"
)

type Repository struct {
	Account  account.Repository
	Example  example.Repository
	Product  product.Repository
	Producer producer.Repository
	Storage  storage.Repository
	Checkout checkout.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Account:  account.NewAccountPG(getClient),
		Example:  example.NewPG(getClient),
		Product:  product.NewPG(getClient),
		Producer: producer.NewPG(getClient),
		Storage:  storage.NewPG(getClient),
		Checkout: checkout.NewPG(getClient),
	}
}
