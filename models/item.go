package models

import (
	"time"
)

type Item struct {
	ItemID int `gorm:"primaryKey;not null;type:int"`
	ItemCode string
	Quantity int
	Description string
	OrderID int 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}