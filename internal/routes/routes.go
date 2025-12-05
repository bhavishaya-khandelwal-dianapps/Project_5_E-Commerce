package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		AuthRoutes(api)
		UserRoutes(api)
		AdminRoutes(api)
		ProductRoutes(api)
		CartRoutes(api)
		OrderRoutes(api)
		PaymentRoutes(api)
	}
}
