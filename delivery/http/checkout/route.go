package checkout

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/store/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("/add-to-cart", r.AddToCart)
	group.GET("/carts", r.GetList)
	group.PUT("/cart/remove", r.RemoveFromCart)

	group.POST("/order", r.CreateCustomerOrder)
}
