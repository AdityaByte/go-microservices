package usecase

import (
	"context"
	"time"

	"github.com/AdityaByte/product-service/internal/domain"
	"github.com/AdityaByte/product-service/internal/repository"
)

func CreateProduct(product *domain.Product) error {
	// Here we have to save the product to the database.
	context, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()
	err := repository.Save(context, product)
	if err != nil {
		return err
	}
	return nil
}

func FindProducts() ([]domain.Product, error) {
	context, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()
	products, err := repository.FindAll(context)
	if err != nil {
		return nil, err
	}
	return products, nil
}
