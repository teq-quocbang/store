package checkout

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	UpsertCart(context.Context, *model.Cart) error
	GetCartByConstraint(ctx context.Context, accountID uuid.UUID, productID uuid.UUID) (model.Cart, error)
	GetListCart(ctx context.Context, accountID uuid.UUID) ([]model.Cart, error)
	RemoveFromCart(ctx context.Context, accountID uuid.UUID, productID uuid.UUID, qty int64) error

	CreateCustomerOrder(context.Context, *model.CustomerOrder) error
	GetListOrdered(
		ctx context.Context,
		accountID uuid.UUID,
		startTime time.Time,
		endTime time.Time,
		order []string) ([]model.CustomerOrder, error)
}
