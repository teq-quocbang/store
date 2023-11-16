package statistics

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	GetProductSoldChart(context.Context, *payload.GetChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error)
	GetProductGrowthChart(context.Context, *payload.GetChartRequest) (*presenter.ListStatisticsSoldProductChartResponseWrapper, error)
}
