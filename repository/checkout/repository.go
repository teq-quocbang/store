package checkout

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	UpsertCart(context.Context, *model.Cart) error
	GetCartByConstraint(ctx context.Context, accountID uuid.UUID, productID uuid.UUID) (model.Cart, error)
	GetListCart(ctx context.Context, accountID uuid.UUID) ([]model.Cart, error)
}
