package domain

import "github.com/google/uuid"

type Order struct {
	Id             uuid.UUID        `gorm:"primarykey" json:"id"`
	OrderNumber    string           `json:"orderNumber"`
	OrderLineItems []OrderLineItems `gorm:"foreignKey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orderLineItems"`
}

type OrderLineItems struct {
	Id       uuid.UUID `gorm:"primaryKey" json:"id"`
	SkuCode  string    `json:"skuCode"`
	Price    float64   `json:"price"`
	Quantity uint16    `json:"quantity"`
}
