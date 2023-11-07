package semester

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/presenter"
)

// Get semester by id
// @Summary Get a semester
// @Description Get semester by id
// @Tags Example
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.SemesterResponseWrapper
// @Router /semester/{id} [get] .
func (r *Route) Get(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = c.Param("id")
		resp *presenter.SemesterResponseWrapper
	)

	resp, err := r.UseCase.Semester.GetByID(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
