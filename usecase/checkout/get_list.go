package checkout

import (
	"context"

	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetListCart(ctx context.Context) (*presenter.ListCartResponseWrapper, error) {
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	carts, err := u.Checkout.GetListCart(ctx, userPrinciple.User.ID)
	if err != nil {
		return nil, myerror.ErrCartGet(err)
	}
	return &presenter.ListCartResponseWrapper{
		Cart: carts,
	}, nil
}
