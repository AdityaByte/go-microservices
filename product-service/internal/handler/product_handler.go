package handler

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/AdityaByte/product-service/internal/domain"
	"github.com/AdityaByte/product-service/internal/usecase"
)

func HandleCreateProduct(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	defer r.Body.Close()

	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		http.Error(w, "Content Type Not Found", http.StatusBadRequest)
		return
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, "Failed to parse the content type", http.StatusBadRequest)
		return
	}

	if mediaType != "application/json" {
		http.Error(w, "Content Not supported", http.StatusBadRequest)
		return
	}

	productData := domain.Product{}

	decoder := json.NewDecoder(body)
	decoder.Decode(&productData)

	if err := usecase.CreateProduct(&productData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product Created Successfully")
}

func HandleFindProducts(w http.ResponseWriter, r *http.Request) {

	products, err := usecase.FindProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}
