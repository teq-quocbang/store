package register

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Update register by id
// @Summary Update an register
// @Description Update register by id
// @Tags Example
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Param req body payload.UnRegisterRequest true "Register info"
// @Success 200 {object} presenter.RegisterResponseWrapper
// @Router /registers/cancel [put] .
func (r *Route) Update(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.UnRegisterRequest{}
		resp *presenter.RegisterResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Register.UnRegister(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
