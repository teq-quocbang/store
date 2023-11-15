package checkout

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/presenter"
)

// GetList cart
// @Summary Get list carts
// @Description Get list carts
// @Tags Cart
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListCartResponseWrapper
// @Router /checkout/carts [get] .
func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		resp *presenter.ListCartResponseWrapper
	)

	resp, err := r.UseCase.Checkout.GetListCart(ctx)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
