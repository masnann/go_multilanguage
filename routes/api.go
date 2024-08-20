package routes

import (
	"multilanguage/handler"
	producthandler "multilanguage/handler/productHandler"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {
	productHandler := producthandler.NewProductHandler(handler)

	e.POST("/products", productHandler.FindListProduct)
	e.POST("/products/create", productHandler.CreateProduct)
}
