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
// @Param req body payload.ListSemesterInformationRequest
// @Success 200 {object} presenter.ListRegisterResponseWrapper
// @Router /register [get] .
func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.ListRegisterInformationRequest{}
		resp *presenter.ListRegisterResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Register.GetListBySemester(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
