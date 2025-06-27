package main

import (
	"log"
	"net/http"

	"github.com/AdityaByte/order-service/internal/handler"
)

func main() {
	http.HandleFunc("/api/order", handler.HandleOrder)
	log.Println("Starting order-service at port :8081..")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Failed to start the server at :8081 %w", err)
	}
}
