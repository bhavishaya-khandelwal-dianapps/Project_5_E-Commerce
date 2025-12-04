package routes

import (
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/handlers"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/products")
	{
		product.POST("/create", middleware.Auth(), middleware.IsAdmin(), handlers.CreateProduct)
		product.GET("/all", middleware.Auth(), handlers.GetAllProducts)
		product.GET("/:id", middleware.Auth(), handlers.GetProduct)
		product.PUT("/:id", middleware.Auth(), middleware.IsAdmin(), handlers.UpdateProduct);
		product.DELETE("/:id", middleware.Auth(), middleware.IsAdmin(), handlers.DeleteProduct);
	}
}
