package routes

import (
	"Engineering/Golang/EcommerceAPI/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	// GET

	incomingRoutes.GET("/users/getUser/:user_id",
		controllers.GetUser())

	incomingRoutes.GET("/users/search",
		controllers.SearchProduct())

	incomingRoutes.GET("/users/productview",
		controllers.ViewProduct())

	// POST
	incomingRoutes.POST("/users/signup",
		controllers.Signup())

	incomingRoutes.POST("/users/login",
		controllers.Login())

	incomingRoutes.POST("/admin/addproduct",
		controllers.AddProduct())

	// PATCH
	incomingRoutes.PATCH("/admin/updateproduct/:product_id",
		controllers.UpdateProduct())

	incomingRoutes.PATCH("/users/updateUser/:user_id",
		controllers.UpdateUser())

	/// DELETE
	incomingRoutes.DELETE("/admin/deleteproduct/:product_id",
		controllers.DeleteProduct())

	incomingRoutes.DELETE("/users/deleteUser/:user_id",
		controllers.DeleteUser())

}
