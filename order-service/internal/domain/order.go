package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id             uint             `gorm:"primarykey;autoIncrement" json:"id"`
	OrderNumber    string           `json:"orderNumber"`
	OrderLineItems []OrderLineItems `gorm:"foreignKey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orderLineItems"`
}

type OrderLineItems struct {
	gorm.Model
	Id       uint    `gorm:"primaryKey" json:"id"`
	SkuCode  string  `json:"skuCode"`
	Price    float64 `json:"price"`
	Quantity uint16  `json:"quantity"`
}
