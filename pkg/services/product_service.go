package services

import (
	"ecomm-crud/pkg/models"
	"ecomm-crud/pkg/repositories"
	"fmt"
)

func CreateProduct(product *models.Product) *models.Product {

	productRepo := repositories.NewProductRepository()

	product, err := productRepo.CreateProduct(product)
	if err != nil {
		fmt.Printf("log.Logger: %v\n", err.Error())
	}
	return product
}
