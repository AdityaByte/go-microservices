package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AdityaByte/inventory-service/internal/usecase"
)

func HandleIsInStock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	skuCodes := r.URL.Query()["skuCode"]

	response, err := usecase.IsInStock(&skuCodes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode the data", http.StatusInternalServerError)
		return
	}
}
