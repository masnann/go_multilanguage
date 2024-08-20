package productservice

import (
	"multilanguage/models"
	"multilanguage/service"
)

type ProductService struct {
	service service.Service
}

func NewProductService(service service.Service) ProductService {
	return ProductService{
		service: service,
	}
}

func (s ProductService) FindListProduct(req models.RequestListProduct) ([]models.ProductModels, error) {
	return s.service.ProductRepo.FindListProduct(req.Language)
}

func (s ProductService) CreateProduct(req models.ProductCreateRequest) (int64, error) {
	product := models.ProductModels{
		Name:        req.Name["en"],
		Description: req.Description["en"],
		Price:       req.Price,
	}

	productID, err := s.service.ProductRepo.CreateProduct(product)
	if err != nil {
		return productID, err
	}

	for language, name := range req.Name {
		if err := s.service.ProductRepo.AddTranslation(models.TranslationCreateRequest{
			EntityType:  "product",
			EntityID:    productID,
			Language:    language,
			FieldName:   "name",
			Translation: name,
		}); err != nil {
			return productID, err
		}
	}

	for language, description := range req.Description {
		if err := s.service.ProductRepo.AddTranslation(models.TranslationCreateRequest{
			EntityType:  "product",
			EntityID:    productID,
			Language:    language,
			FieldName:   "description",
			Translation: description,
		}); err != nil {
			return productID, err
		}
	}
	return productID, nil
}
