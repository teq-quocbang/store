package statistics

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

// GetList statistics product sold chart
// @Summary Get statistics product sold chart
// @Description Get statistics product sold chart
// @Tags Statistics
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListCartResponseWrapper
// @Router /statistics/product-sold-chart [get] .
func (r *Route) GetProductSoldChart(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetChartRequest{}
		resp *presenter.ListStatisticsSoldProductChartResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Statistics.GetProductSoldChart(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

// GetList statistics product sold chart
// @Summary Get statistics product sold chart
// @Description Get statistics product sold chart
// @Tags Statistics
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListCartResponseWrapper
// @Router /statistics/product-sold-chart [get] .
func (r *Route) GetProductGrowthChart(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetChartRequest{}
		resp *presenter.ListStatisticsSoldProductChartResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Statistics.GetProductGrowthChart(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
