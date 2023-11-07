package example

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
	group.GET("/:id", r.GetByID)
	group.PUT("/:id", r.Update)
	group.DELETE("/:id", r.Delete)
}
