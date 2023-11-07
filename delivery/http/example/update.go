package example

import (
	"strconv"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Update example by id
// @Summary Update an example
// @Description Update example by id
// @Tags Example
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Param req body payload.UpdateExampleRequest true "Example info"
// @Success 200 {object} presenter.ExampleResponseWrapper
// @Router /examples/{id} [put] .
func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.ExampleResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	req := payload.UpdateExampleRequest{
		ID: id,
	}

	if err = c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.Example.Update(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
