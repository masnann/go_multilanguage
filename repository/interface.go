package repository

import "multilanguage/models"

type ProductRepositoryInterface interface {
	FindListProduct(language string) ([]models.ProductModels, error)
	CreateProduct(req models.ProductModels) (int64, error)
	AddTranslation(req models.TranslationCreateRequest) error
}
