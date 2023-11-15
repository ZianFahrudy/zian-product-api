package service

import (
	"context"
	"zian-product-api/domain/entity"
	"zian-product-api/domain/repository"
)

type ProductService interface {
	GetProductList(ctx context.Context) ([]entity.Product, error)
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{ProductRepository: productRepository}
}

func (s *productServiceImpl) GetProductList(ctx context.Context) ([]entity.Product, error) {
	products, err := s.ProductRepository.FindAll(ctx)

	if err != nil {
		return products, err
	}

	return products, nil
}
