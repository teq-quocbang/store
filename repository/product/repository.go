package product

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Product) error
}