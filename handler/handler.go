package handler

import "multilanguage/service"

type Handler struct {
	ProductService service.ProductServiceInterface
}

func NewHandler(
	ProductService service.ProductServiceInterface,
) Handler {
	return Handler{
		ProductService: ProductService,
	}
}
