package semester

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Create semester
// @Summary Create semester
// @Description create a semester
// @Tags Semester
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.CreateSemesterRequest
// @Success 200 {object} presenter.SemesterResponseWrapper
// @Router /user/sign-up [post] .
func (r *Route) CreateSemester(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.CreateSemesterRequest{}
		resp *presenter.SemesterResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Semester.CreateSemester(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
