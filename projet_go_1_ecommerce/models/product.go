package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Quantity    int      `json:"quantity"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category"`
}
