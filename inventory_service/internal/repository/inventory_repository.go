package repository

import (
	"log"

	"github.com/AdityaByte/inventory-service/internal/domain"
	"github.com/AdityaByte/inventory-service/internal/dto"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var database *gorm.DB

func init() {
	dsn := "root:aditya@tcp(localhost:3306)/inventory_service1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// Here we have to clear the db if any entry is prexisted in the database.
	inventories := [2]domain.Inventory{
		{
			Id:       uuid.New(),
			SkuCode:  "iphone12",
			Quantity: 100,
		},
		{
			Id:       uuid.New(),
			SkuCode:  "iphone12-blue",
			Quantity: 0,
		},
	}

	db.AutoMigrate(&domain.Inventory{}) // Create the table if not existed in the db.

	// deleting the items if existed previously just for testing purpose.
	db.Delete(&inventories)
	// Inserting data to db.
	db.Create(&inventories)

	database = db
	log.Println("Data inserted successfully to db")
}

func FindSkuCodeIn(skuCodes *[]string) ([]dto.InventoryResponse, error) {

	var inventoryResponses []dto.InventoryResponse
	for _, value := range *skuCodes {
		var inventory domain.Inventory
		database.Find(&inventory, "sku_code = ?", value)
		if inventory == (domain.Inventory{}) {
			inventoryResponses = append(inventoryResponses, dto.InventoryResponse{
				SkuCode: value,
				IsInStock: false,
			})
		} else {
			if inventory.Quantity > 0 {
				inventoryResponses = append(inventoryResponses, dto.InventoryResponse{
					SkuCode: value,
					IsInStock: true,
				})
			} else {
				inventoryResponses = append(inventoryResponses, dto.InventoryResponse{
					SkuCode: value,
					IsInStock: false,
				})
			}
		}
	}

	return inventoryResponses, nil
}
