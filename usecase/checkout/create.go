package checkout

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) AddToCard(ctx context.Context, req *payload.AddToCartRequest) (*presenter.CartResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCartInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return nil, myerror.ErrCartInvalidParam(err.Error())
	}

	cart := &model.Cart{
		AccountID: userPrinciple.User.ID,
		ProductID: productID,
		Qty:       req.Qty,
	}
	err = u.Checkout.UpsertCart(ctx, cart)
	if err != nil {
		return nil, myerror.ErrCartCreate(err)
	}

	return &presenter.CartResponseWrapper{
		Cart: cart,
	}, nil
}
