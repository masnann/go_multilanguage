package app

import (
	"multilanguage/handler"
	"multilanguage/repository"
	productrepository "multilanguage/repository/productRepository"
	"multilanguage/service"
	productservice "multilanguage/service/productService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	productRepo := productrepository.NewProductRepository(repo)

	service := service.NewService(productRepo)

	productService := productservice.NewProductService(service)

	handler := handler.NewHandler(productService)

	return handler
}
