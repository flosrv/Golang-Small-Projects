package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNumber string      `json:"order_number" gorm:"uniqueIndex"`
	OrderDate   string      `json:"order_date"`
	OrderStatus string      `json:"order_status"`
	OrderItems  []OrderItem `json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	// Product Product `json:"product"`
}
