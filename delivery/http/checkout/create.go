package checkout

import (
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Upsert
// @Summary Upsert cart
// @Description create a cart
// @Tags Storage
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.AddToCartRequest
// @Success 200 {object} presenter.CartResponseWrapper
// @Router /checkout/add-to-cart [post] .
func (r *Route) AddToCart(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.AddToCartRequest{}
		resp *presenter.CartResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Checkout.AddToCard(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

// Create
// @Summary Create customer order
// @Description create customer order
// @Tags Checkout
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.AddToCartRequest
// @Success 200 {object} presenter.CartResponseWrapper
// @Router /checkout/order [post] .
func (r *Route) CreateCustomerOrder(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.CustomerOrderRequest{}
		resp *presenter.CustomerOrderResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Checkout.CreateCustomerOrder(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
