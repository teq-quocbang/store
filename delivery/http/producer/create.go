package producer

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// Create
// @Summary Create producer
// @Description create a producer
// @Tags Producer
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.CreateProducerRequest
// @Success 200 {object} presenter.ProducerResponseWrapper
// @Router /producer [post] .
func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = &payload.CreateProducerRequest{}
		resp *presenter.ProducerResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Producer.Create(ctx, req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
