package domain

import "github.com/google/uuid"

type Inventory struct {
	Id       uuid.UUID `gorm:"primarykey" json:"id"`
	SkuCode  string    `json:"skuCode"`
	Quantity uint16    `json:"quantity"`
}

func (Inventory) TableName() string {
	return "inventory"
}