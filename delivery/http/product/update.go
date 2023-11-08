package product

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Update product by id
// @Summary Update an product
// @Description Update product by id
// @Tags Product
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Param req body payload.UpdateProductRequest true "Product info"
// @Success 200 {object} presenter.ProductResponseWrapper
// @Router /product/{id} [put] .
func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		req   = &payload.UpdateProductRequest{}
		resp  *presenter.ProductResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	req.ID = idStr

	resp, err := r.UseCase.Product.Update(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
