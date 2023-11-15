package checkout

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Update cart
// @Summary Update
// @Description Update
// @Tags Checkout
// @Accept json
// @Produce json
// @Security AuthToken
// @Param req body payload.RemoveFormCartRequest true "Product info"
// @Success 200 {object} presenter.ProductResponseWrapper
// @Router /checkout/cart/remove [put] .
func (r *Route) RemoveFromCart(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.RemoveFormCartRequest{}
		resp *presenter.CartResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Checkout.RemoveFromCart(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
