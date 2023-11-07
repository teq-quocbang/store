package course

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
)

// Delete course by id
// @Summary Delete an course
// @Description Delete course by id
// @Tags Course
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200
// @Router /course/{id} [delete] .
func (r *Route) Delete(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
	)

	err := r.UseCase.Course.Delete(ctx, idStr)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, nil)
}
