package models

import "time"

type Order struct {
	OrderID      int `gorm:"primaryKey;not null;type:int"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Item `gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
