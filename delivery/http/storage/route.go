package storage

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("", r.Upsert)
	group.GET("/:locat", r.Upsert)
}
