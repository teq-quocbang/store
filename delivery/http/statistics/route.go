package statistics

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.GET("/product-sold-chart", r.GetProductSoldChart)
	group.GET("/product-growth-chart", r.GetProductGrowthChart)
}
