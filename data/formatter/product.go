package formatter

import "zian-product-api/domain/entity"

type ProductFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func FormatProduct(product entity.Product) ProductFormatter {
	productFormatter := ProductFormatter{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Category:    product.Category,
		Image:       product.Image,
		Description: product.Description,
	}

	return productFormatter
}

func FormatProductList(products []entity.Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)

		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
