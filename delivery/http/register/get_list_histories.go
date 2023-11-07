package register

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// List register
// @Summary list register
// @Description list registers
// @Tags Register
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.ListRegisteredHistories
// @Success 200 {object} presenter.ListRegisterResponseWrapper
// @Router /register/histories [get] .
func (r *Route) GetHistories(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.ListRegisteredHistories{}
		resp *presenter.ListRegisterResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Register.GetListRegisteredHistories(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
