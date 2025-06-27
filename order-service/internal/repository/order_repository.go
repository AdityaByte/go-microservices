package repository

import (
	"log"

	"github.com/AdityaByte/order-service/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	dsn := "root:aditya@tcp(localhost:3306)/order_service1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open the connection to the database %w", err)
	}

	// Auto-migrating
	// Creating the tables if not exists initially
	if err := db.AutoMigrate(&domain.Order{}, &domain.OrderLineItems{}); err != nil {
		log.Fatal("Failed to migrate the database schema %w", err)
	}
	database = db
}

func Save(order *domain.Order) error {
	result := database.Create(order)
	return result.Error
}
