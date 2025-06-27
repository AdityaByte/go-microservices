package handler

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/AdityaByte/order-service/internal/dto"
	"github.com/AdityaByte/order-service/internal/usecase"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := r.Body
	defer r.Body.Close()

	// Validating the content
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		http.Error(w, "Content Type not found", http.StatusBadRequest)
		return
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, "Failed to parse the content", http.StatusBadRequest)
		return
	}

	if mediaType != "application/json" {
		http.Error(w, "Content not supported", http.StatusBadRequest)
		return
	}

	// If we gets the valid content type we have to parse it to the json.
	var orderRequest dto.OrderRequest

	if err := json.NewDecoder(body).Decode(&orderRequest); err != nil {
		http.Error(w, "Failed to parse the body", http.StatusBadRequest)
		return
	}

	// Calling the usecase handler for buisness logic.
	if err := usecase.PlaceOrder(&orderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Order Placed successfully")
}
