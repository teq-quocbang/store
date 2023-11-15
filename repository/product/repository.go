package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Product) error
	CreateList(context.Context, []model.Product) error
	GetList(context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.Product, error)
	Update(context.Context, *model.Product) error
	Delete(context.Context, uuid.UUID) error
	GetListByProductIDs(ctx context.Context, productIDs []uuid.UUID) ([]model.Product, error)
}
