package product

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
	group.PUT("/:id", r.Update)
	group.DELETE("/:id", r.Delete)
	group.POST("/export", r.Export)
}

func ProductsInit(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("", r.CreateList)
	group.POST("/import", r.CreateListWithImportFile)
	group.POST("/thirty-part", r.CreateListWithThirtyPart)
}
