package service

import (
	"product-service/internal/model"
	"product-service/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(product model.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.Repo.GetAllProducts()
}
