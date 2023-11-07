package course

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/presenter"
)

// Get course by id
// @Summary Get a course
// @Description Get course by id
// @Tags Course
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.CourseResponseWrapper
// @Router /course/{id} [get] .
func (r *Route) Get(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = c.Param("id")
		resp *presenter.CourseResponseWrapper
	)

	resp, err := r.UseCase.Course.GetByID(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
