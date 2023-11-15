package repository

import (
	"context"
	"zian-product-api/domain/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *productRepositoryImpl {
	return &productRepositoryImpl{db}
}

func (r *productRepositoryImpl) FindAll(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.WithContext(ctx).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}
