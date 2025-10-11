package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
}
type CartItem struct {
	gorm.Model
	CartId    uint   `json:"cart_id"`
	ProductId uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
}
