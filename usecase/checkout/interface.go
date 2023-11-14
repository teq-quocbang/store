package checkout

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	AddToCard(context.Context, *payload.AddToCartRequest) (*presenter.CartResponseWrapper, error)
	GetListCart(context.Context) (*presenter.ListCartResponseWrapper, error)
}
