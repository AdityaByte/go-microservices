package usecase

import (
	"github.com/AdityaByte/order-service/internal/domain"
	"github.com/AdityaByte/order-service/internal/dto"
	"github.com/AdityaByte/order-service/internal/repository"
	"github.com/google/uuid"
)

func PlaceOrder(orderRequest *dto.OrderRequest) error {
	var order domain.Order
	len := len(orderRequest.OrderLineItemsDtoList)
	for i := 0; i < len; i++ {
		// here we have to fetch the items one by one and
		orderListItem := orderRequest.OrderLineItemsDtoList[i]
		newOrderListItem := domain.OrderLineItems{
			Id:       uuid.New(),
			SkuCode:  orderListItem.SkuCode,
			Price:    orderListItem.Price,
			Quantity: orderListItem.Quantity,
		}
		// Here we have to map those
		order.OrderLineItems = append(order.OrderLineItems, newOrderListItem)
	}

	// Now we have to place the order.
	err := repository.Save(&order)
	if err != nil {
		return err
	}
	return nil
}
