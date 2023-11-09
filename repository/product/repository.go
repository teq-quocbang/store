package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Product) error
	Update(context.Context, *model.Product) error
	Delete(context.Context, uuid.UUID) error
}
