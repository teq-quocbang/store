package register

import (
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
)

// Tracing insufficient credits
// @Summary insufficient credits
// @Description insufficient credits
// @Tags Register
// @Accept  json
// @Produce json
// @Security no
// @Router /register/tracing/insufficient-credits [get] .
func (r *Route) Tracing(c echo.Context) error {
	var (
		ctx = &teq.CustomEchoContext{Context: c}
	)

	err := r.UseCase.Register.TracingInsufficientCreditsStatistics(ctx)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return ctx.NoContent(http.StatusOK)
}
