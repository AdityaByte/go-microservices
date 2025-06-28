package usecase

import (
	"fmt"

	"github.com/AdityaByte/inventory-service/internal/dto"
	"github.com/AdityaByte/inventory-service/internal/repository"
)

func IsInStock(skuCodes *[]string) ([]dto.InventoryResponse, error) {
	response, err := repository.FindSkuCodeIn(skuCodes)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch the reponse from database %w", err)
	}
	return response, nil
}
