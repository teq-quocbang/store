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

	// check storage
	inventoryQty, err := u.Storage.GetInventoryQty(ctx, productID)
	if err != nil {
		return nil, myerror.ErrStorageGet(err)
	}
	if inventoryQty < int(req.Qty) {
		return nil, myerror.ErrStorageInvalidParam("request qty out of the inventory qty")
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

func (u *UseCase) CreateCustomerOrder(ctx context.Context, req *payload.CustomerOrderRequest) (*presenter.CustomerOrderResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCustomerOrderInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return nil, myerror.ErrCustomerOrderInvalidParam(err.Error())
	}

	// get inventory qty
	inventoryQty, err := u.Storage.GetInventoryQty(ctx, productID)
	if err != nil {
		return nil, myerror.ErrStorageGet(err)
	}

	// get product from cart
	cart, err := u.Checkout.GetCartByConstraint(ctx, userPrinciple.User.ID, productID)
	if err != nil {
		return nil, myerror.ErrCartGet(err)
	}
	if inventoryQty < int(cart.Qty) {
		return nil, myerror.ErrCustomerOrderInvalidParam("request qty in cart is out of inventory qty")
	}
	if cart.Qty < 1 {
		return nil, myerror.ErrCustomerOrderInvalidParam("cart is empty")
	}

	// get price from product
	product, err := u.Product.GetByID(ctx, cart.ProductID)
	if err != nil {
		return nil, myerror.ErrProductGet(err)
	}

	customerOrder := &model.CustomerOrder{
		AccountID:  userPrinciple.User.ID,
		ProductID:  productID,
		PriceOfPer: product.Price,
		SoldQty:    cart.Qty,
	}
	if err := u.Checkout.CreateCustomerOrder(ctx, customerOrder); err != nil {
		return nil, myerror.ErrCustomerOrderCreate(err)
	}

	return &presenter.CustomerOrderResponseWrapper{
		CustomerOrder: customerOrder,
	}, nil
}
