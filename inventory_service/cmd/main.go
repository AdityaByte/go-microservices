package main

import (
	"log"
	"net/http"

	"github.com/AdityaByte/inventory-service/internal/handler"
)

func main() {
	http.HandleFunc("/api/inventory", handler.HandleIsInStock)
	log.Println("Starting inventory-service at port :8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("Failed to start the inventory-service at :8082, %w", err)
	}
}
