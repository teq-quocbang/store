package course

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Create course
// @Summary Create course
// @Description create a course
// @Tags Course
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.CreateCourseRequest
// @Success 200 {object} presenter.CourseResponseWrapper
// @Router /course [post] .
func (r *Route) CreateCourse(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.CreateCourseRequest{}
		resp *presenter.CourseResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Course.CreateCourse(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
