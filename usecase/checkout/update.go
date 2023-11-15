package checkout

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) RemoveFromCart(ctx context.Context, req *payload.RemoveFormCartRequest) (*presenter.CartResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCartInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return nil, myerror.ErrCartInvalidParam(err.Error())
	}

	cartModel, err := u.Checkout.GetCartByConstraint(ctx, userPrinciple.User.ID, productID)
	if err != nil {
		return nil, myerror.ErrCartGet(err)
	}
	if cartModel.Qty < req.Qty {
		req.Qty = cartModel.Qty // request qty out of cart qty then swap it
	}
	if cartModel.Qty < 1 {
		return nil, myerror.ErrCartInvalidParam("card was cleared")
	}

	err = u.Checkout.RemoveFromCart(ctx, userPrinciple.User.ID, productID, req.Qty)
	if err != nil {
		return nil, myerror.ErrCartUpdate(err)
	}

	cartModel.Qty = cartModel.Qty - req.Qty // set qty after update
	return &presenter.CartResponseWrapper{
		Cart: &cartModel,
	}, nil
}
