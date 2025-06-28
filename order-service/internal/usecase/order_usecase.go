package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

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

	var wg sync.WaitGroup
	errorChan := make(chan error, 1)

	wg.Add(1)
	go func() {
		req, err := http.NewRequest("GET", "http://localhost:8082/api/inventory", nil)
		if err != nil {
			errorChan <- err
			wg.Done()
			return
		}

		queryParams := url.Values{}
		for _, value := range order.OrderLineItems {
			queryParams.Add("skuCode", value.SkuCode)
		}

		req.URL.RawQuery = queryParams.Encode()

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- err
		}
		defer resp.Body.Close()

		var inventoryResponses []dto.InventoryResponse

		if err := json.NewDecoder(resp.Body).Decode(&inventoryResponses); err != nil {
			errorChan <- err
			wg.Done()
			return
		}

		fmt.Println(inventoryResponses)
		log.Println("Decoded the data successfully...")

		var isInStock bool
		for _, inventoryResponse := range inventoryResponses {
			if !inventoryResponse.IsInStock {
				isInStock = false
				break
			} else {
				isInStock = true
			}
		}

		if isInStock {
			errorChan <- nil
			log.Println("Order placed successfully")
		} else {
			errorChan <- fmt.Errorf("Item is not in the stock")
		}

		wg.Done()
	}()

	wg.Wait()

	select {
	case err := <-errorChan:
		if err != nil {
			return err
		} else {
			repository.Save(&order)
			return nil
		}
	default:
		repository.Save(&order)
		return nil
	}

}
