package register

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("", r.Create)
	group.GET("", r.GetList)
	group.GET("/histories", r.GetHistories)
	group.GET("/tracing/insufficient-credits", r.Tracing)
	group.PUT("/cancel", r.Update)
}
