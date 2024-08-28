package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ghssni/Rentalind-Go-App/handlers"
	"github.com/ghssni/Rentalind-Go-App/middleware"
	"github.com/ghssni/Rentalind-Go-App/utils"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine) {
	// Initialize database connection
	err := utils.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	// Define API routes
	api := router.Group("/api")

	// User routes
	api.POST("/users", handlers.CreateUserHandler)
	api.GET("/users/:id", handlers.GetUserHandler)

	// Game routes
	api.POST("/games", handlers.CreateGameHandler)
	api.GET("/games/:id", handlers.GetGameHandler)

	// Rental routes
	api.POST("/rentals", handlers.CreateRentalHandler)
	api.GET("/rentals/:id", handlers.GetRentalHandler)

	// Apply authentication middleware to protected routes
	// Example:
	// protected := api.Group("/protected")
	// protected.Use(middleware.AuthMiddleware)
	// protected.GET("/some-route", someHandler)
}
