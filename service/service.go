package service

import "multilanguage/repository"

type Service struct {
	ProductRepo repository.ProductRepositoryInterface
}

func NewService(
	ProductRepo repository.ProductRepositoryInterface,
) Service {
	return Service{
		ProductRepo: ProductRepo,
	}
}
