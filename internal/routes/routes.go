package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		// User routes
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Logic) // login will be implemented next
		// api.GET("/users/:id", handlers.GetUser) // get user by ID

		// // Product routes
		// api.POST("/products", handlers.CreateProduct)
		// api.GET("/products", handlers.GetProducts)
		// api.GET("/products/:id", handlers.GetProduct)
		// api.PUT("/products/:id", handlers.UpdateProduct)
		// api.DELETE("/products/:id", handlers.DeleteProduct)

		// // TODO: Add routes for Cart, Order, Review, Payment
	}
}
