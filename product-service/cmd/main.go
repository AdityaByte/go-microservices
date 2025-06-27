package main

import (
	"log"
	"net/http"

	"github.com/AdityaByte/product-service/internal/handler"
)

func main() {
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.HandleCreateProduct(w, r)
		case http.MethodGet:
			handler.HandleFindProducts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	log.Println("Starting server at port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start the server at port :8080")
	}
}
