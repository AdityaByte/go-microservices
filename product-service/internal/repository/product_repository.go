package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/AdityaByte/product-service/internal/config"
	"github.com/AdityaByte/product-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func FindAll(ctx context.Context) ([]domain.Product, error) {

	var data []domain.Product

	findOptions := options.Find()
	cur, err := mongoRepository.Collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve all documents from database %w", err)
	}

	for cur.Next(ctx) {
		var element domain.Product
		err := cur.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, element)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)

	return data, nil
}