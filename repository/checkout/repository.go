package checkout

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	UpsertCart(context.Context, *model.Cart) error
}
