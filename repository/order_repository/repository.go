package order_repository

import "assignment2/models"

type Repository interface {
	// SQL
	// CreateOrder(orderPayload models.Order, itemPayload []models.Item) error

	//GORM
	CreateOrder(orderPayload *models.Order) error
	GetOrders() (*[]models.Order,error)
	GetOrder(id int) (*models.Order, error)
	UpdateOrder(orderPayload *models.Order) (*models.Order, error)
	DeleteOrder(id int) error
}
