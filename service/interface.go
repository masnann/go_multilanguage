package service

import "multilanguage/models"

type ProductServiceInterface interface {
	FindListProduct(req models.RequestListProduct) ([]models.ProductModels, error)
	CreateProduct(req models.ProductCreateRequest) (int64, error)
}
