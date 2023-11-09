package product

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Create
// @Summary Create product
// @Description create a product
// @Tags Product
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.CreateProductRequest
// @Success 200 {object} presenter.ProductResponseWrapper
// @Router /product [post] .
func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.CreateProductRequest{}
		resp *presenter.ProductResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Product.Create(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
