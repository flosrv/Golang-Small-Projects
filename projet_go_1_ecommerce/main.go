package main

import (
	"go_learn_ecommerce/config"
	"go_learn_ecommerce/models"
)

func main() {

	config.ConnectDB()
	err := config.DB.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Order{},
		&models.OrderItem{},
		&models.Cart{},
		&models.CartItem{},
	)
	if err != nil {
		panic(err)
	}

	// config.DB.Create(&models.Product{Title: "Product 1", Price: 100})
}
