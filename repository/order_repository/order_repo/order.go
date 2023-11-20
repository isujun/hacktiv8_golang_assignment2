package order_repo

import (
	"assignment2/models"
	"assignment2/repository/order_repository"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) order_repository.Repository {
	return &orderRepo{db: db}
}

func (orderRepo *orderRepo) CreateOrder(orderPayload *models.Order) error {
	var Orders = *orderPayload
	err := orderRepo.db.Create(&Orders).Error
	return err
}

func (orderRepo *orderRepo) GetOrders() (*[]models.Order, error) {
	var Orders []models.Order
	err := orderRepo.db.Preload("Items").Find(&Orders).Error

	if err != nil {
		return nil, err
	}

	return &Orders, nil
}

func (orderRepo *orderRepo) GetOrder(id int) (*models.Order, error) {
	var Order models.Order
	err := orderRepo.db.Preload("Items").First(&Order, "order_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &Order, nil
}

func (orderRepo *orderRepo) UpdateOrder(orderPayload *models.Order) (*models.Order, error) {
	var Order = *orderPayload
	items := Order.Items
	var failed = false
	var err error

	orderRepo.db.Begin()
	for _, item := range items {
		err = orderRepo.db.Model(&models.Item{}).Where("item_code = ?", item.ItemCode).Updates(models.Item{Description: item.Description, Quantity: item.Quantity}).Error
		if err != nil {
			failed = true
			break
		}
	}
	if failed {
		orderRepo.db.Rollback()
		return nil, err
	}

	err = orderRepo.db.Model(&Order).Updates(models.Order{CreatedAt: Order.CreatedAt, CustomerName: Order.CustomerName}).Error
	if err != nil {
		orderRepo.db.Rollback()
		return nil, err
	}
	orderRepo.db.Commit()
	return &Order, nil
}

func (orderRepo *orderRepo) DeleteOrder(orderId int) error {
	err := orderRepo.db.Where("order_id = ?", orderId).Delete(&models.Item{}).Error
	if err != nil {
		return err
	}
	err = orderRepo.db.Where("order_id = ?", orderId).Delete(&models.Order{}).Error
	return err
}
