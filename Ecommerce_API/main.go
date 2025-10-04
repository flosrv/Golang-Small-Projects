package main

import (
	"Ecommerce_API/controllers"
	"Ecommerce_API/database"
	"Ecommerce_API/middleware"
	"Ecommerce_API/models"
	"Ecommerce_API/routes"
	"Ecommerce_API/tokens"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func main (){

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(
		database.Client, "Products")
	, database.UserData(database.Client, "Users")) 
		}
