package dto

type OrderRequest struct {
	OrderLineItemsDtoList []OrderLineItemsDto `json:"orderLineItemsDtoList"`
}

type OrderLineItemsDto struct {
	Id       uint    `json:"id"`
	SkuCode  string  `json:"skuCode"`
	Price    float64 `json:"price"`
	Quantity uint16  `json:"quantity"`
}
