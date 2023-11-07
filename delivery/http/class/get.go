package class

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/presenter"
)

// Get class by id
// @Summary Get a class
// @Description Get class by id
// @Tags Class
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.ClassResponseWrapper
// @Router /class/{id} [get] .
func (r *Route) Get(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = c.Param("id")
		resp *presenter.ClassResponseWrapper
	)

	resp, err := r.UseCase.Class.GetByID(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
