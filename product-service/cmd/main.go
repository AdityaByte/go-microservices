package main

import (
	"log"
	"net/http"

	"github.com/AdityaByte/product-service/internal/handler"
)

func main() {
	http.HandleFunc("/api/product", handler.HandleCreateProduct)
	log.Println("Starting server at port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start the server at port :8080")
	}
}
