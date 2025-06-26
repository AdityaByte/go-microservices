package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/AdityaByte/product-service/internal/config"
	"github.com/AdityaByte/product-service/internal/domain"
)

var mongoRepository config.MongoRepository

func init()  {
	repo, err := config.LoadMongoRepository()
	if err != nil {
		log.Panic(err.Error())
		return;
	}
	mongoRepository = *repo
}

func Save(ctx context.Context, product *domain.Product) error {
	_, err := mongoRepository.Collection.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("Failed to insert the document to database %w", err)
	}
	return nil
}