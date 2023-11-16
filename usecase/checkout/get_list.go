package checkout

import (
	"context"

	"github.com/shopspring/decimal"

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

	cartResponseInformation := make([]presenter.CartInformation, len(carts))
	for i, cart := range carts {
		product, err := u.Product.GetByID(ctx, cart.ProductID)
		if err != nil {
			return nil, myerror.ErrProducerGet(err)
		}
		totalPrice := product.Price.IntPart() * cart.Qty
		cartResponseInformation[i] = presenter.CartInformation{
			Cart:       cart,
			TotalPrice: decimal.NewFromInt32(int32(totalPrice)),
		}
	}

	return &presenter.ListCartResponseWrapper{
		Cart: cartResponseInformation,
	}, nil
}
