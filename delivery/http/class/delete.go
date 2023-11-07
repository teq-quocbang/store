package class

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
)

// Delete class by id
// @Summary Delete an class
// @Description Delete class by id
// @Tags Class
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200
// @Router /class/{id} [delete] .
func (r *Route) Delete(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
	)

	err := r.UseCase.Class.Delete(ctx, idStr)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, nil)
}
