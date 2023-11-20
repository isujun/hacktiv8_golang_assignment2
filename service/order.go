package service

import (
	"assignment2/dto"
	"assignment2/dto/request_response"
	"assignment2/models"
	"assignment2/repository/order_repository"
	"strconv"

)

type orderService struct {
	OrderRepo order_repository.Repository
}

type OrderService interface {
	CreateOrder(newOrderRequest *dto.NewOrderRequest) *request_response.Response
	GetOrders() *request_response.Response
	UpdateOrder(orderId int, newOrderRequest *dto.NewOrderRequest) *request_response.Response
	DeleteOrder(orderId int) *request_response.Response
}

func NewOrderService(orderRepo order_repository.Repository) OrderService {
	return &orderService{OrderRepo: orderRepo}
}

// GORM
func (os *orderService) CreateOrder(newOrderRequest *dto.NewOrderRequest) *request_response.Response {
	itemPayload := []models.Item{}
	orderPayload := models.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
		Items:        itemPayload,
	}
	for _, eachItem := range newOrderRequest.Items {
		item := models.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}
		itemPayload = append(itemPayload, item)
	}
	orderPayload.Items = append(orderPayload.Items, itemPayload...)

	err := os.OrderRepo.CreateOrder(&orderPayload)
	if err != nil {
		return request_response.BadRequestResponse(err)
	}

	return request_response.SuccessCreateResponse(orderPayload)
}

func (os *orderService) GetOrders() *request_response.Response {
	orders, err := os.OrderRepo.GetOrders()
	if err != nil {
		return request_response.InternalServerErrorResponse(err)
	}
	return request_response.SuccessFindResponse(orders)
}

func (os *orderService) UpdateOrder(orderId int, newOrderRequest *dto.NewOrderRequest) *request_response.Response {
	order, err := os.OrderRepo.GetOrder(orderId)
	if err != nil {
		msg := "order with id:" + strconv.Itoa(orderId) + "not found"
		return request_response.NotFoundResponse(msg)
	}

	itemPayload := []models.Item{}
	orderPayload := models.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
		Items:        itemPayload,
	}
	for _, eachItem := range newOrderRequest.Items {
		item := models.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}
		itemPayload = append(itemPayload, item)
	}
	orderPayload.Items = itemPayload

	order.OrderedAt = orderPayload.OrderedAt
	order.CustomerName = orderPayload.CustomerName

	for _, eachItemPayload := range orderPayload.Items {
		isFound := false

		for _, eachItem := range order.Items{
			if eachItemPayload.ItemCode == eachItem.ItemCode {
				isFound = true
				break
			}
		}
		if !isFound {
			msg := "item with code: " + eachItemPayload.ItemCode + " is not found"
			return request_response.NotFoundResponse(msg)
		}
	}
	order.Items = orderPayload.Items

	updatedOrder, err := os.OrderRepo.UpdateOrder(order)
	if err != nil {
		return request_response.InternalServerErrorResponse(err)
	}
	return request_response.SuccessUpdateResponse(updatedOrder)
}

func (os *orderService) DeleteOrder(orderId int) *request_response.Response {
	_, err := os.OrderRepo.GetOrder(orderId)
	if err != nil {
		msg := "order with id:" + strconv.Itoa(orderId) + "not found"
		return request_response.NotFoundResponse(msg)
	}

	err = os.OrderRepo.DeleteOrder(orderId)
	if err != nil {
		return request_response.InternalServerErrorResponse(err)
	}
	msg := "order with id: " + strconv.Itoa(orderId) + " successfully deleted"
	return request_response.SuccessDeleteResponse(msg)
}

// // SQL
// func (os *orderService) CreateOrder(newOrderRequest dto.NewOrderRequest) error {
// 	orderPayload := models.Order{
// 		OrderedAt:    newOrderRequest.OrderedAt,
// 		CustomerName: newOrderRequest.CustomerName,
// 	}

// 	itemPayload := []models.Item{}

// 	for _, eachItem := range newOrderRequest.Items {
// 		item := models.Item{
// 			ItemCode: eachItem.ItemCode,
// 			Description: eachItem.Description,
// 			Quantity: eachItem.Quantity,
// 		}
// 		itemPayload = append(itemPayload, item)
// 	}

// 	err := os.OrderRepo.CreateOrder(orderPayload, itemPayload)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
