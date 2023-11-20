package order_pg

import (
	"assignment2/models"
	"assignment2/repository/order_repository"
	"database/sql"
	"fmt"
)

type orderPG struct {
	db *sql.DB
}

const (
	createOrderQuery = `
	INSERT INTO "orders" ("ordered_at", "customer_name")
	VALUES (:ordered_at, :customer_name)
	RETURNING "order_id"
	`

	createItemQuery = `
	INSERT INTO "items" ("item_code", "description", "quantity", "order_id")
	VALUES (:item_code, :description, :quantity, :order_id)
	`
)

func NewOrderPG(db *sql.DB) order_repository.Repository {
	return &orderPG{db: db}
}

func (o *orderPG) CreateOrder(order models.Order, items []models.Item) error {
	tx, err := o.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	var orderID int

	orderRow := tx.QueryRow(createOrderQuery, order.OrderedAt, order.CustomerName)
	
	err = orderRow.Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to scan order row: %v", err)
	}

	for _, item := range items {
		_, err := tx.Exec(createItemQuery, item.ItemCode, item.Description, item.Quantity, orderID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert item: %v", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}
