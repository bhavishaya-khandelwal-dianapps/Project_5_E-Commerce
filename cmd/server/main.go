package main

import (
	"log"
	"os"
	"time"

	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/E-Commerce-Website/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode for production
	// Gin has three modes: debug (default), release, and test.
	// Release mode: hides debug logs and is faster for production.
	// Equivalent to NODE_ENV=production in Node.js.
	gin.SetMode(gin.TestMode)

	// Initialize config and database
	config.Init()

	// Auto-migrate all models
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
		&models.Review{},
		&models.Payment{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	// Create Gin router
	router := gin.New()
	// gin.New() creates a new HTTP server without any default middleware.
	// Alternative: gin.Default() creates a server with logger and recovery middleware by default.
	// Here, you manually attach middleware next.
	// gin.Logger() → logs every HTTP request (method, path, status, time).
	// gin.Recovery() → recovers from panics and returns HTTP 500 instead of crashing the server.

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Trusted proxies (replace with your proxy IP or leave empty for localhost)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // frontend URL(s)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// CORS allows your frontend (React/Vue/Angular) to call the backend safely.
	// AllowOrigins → list of frontend URLs allowed to make requests.
	// AllowMethods → HTTP methods allowed.
	// AllowHeaders → which headers frontend can send.
	// ExposeHeaders → which headers backend exposes to frontend.
	// AllowCredentials → allows cookies and auth headers.
	// MaxAge → how long the browser can cache preflight response.

	// Routes
	routes.SetupRoutes(router)
	// Calls your SetupRoutes function.
	// Registers all API endpoints (/api/v1/register, /products, /orders, etc.).

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
