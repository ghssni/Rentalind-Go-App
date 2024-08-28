package routes

import (
	"github.com/gin-gonic/gin"
	
	"Rentalind-Go-App/controllers"

	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes mengatur routing untuk aplikasi dan menghubungkan middleware JWT
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Middleware untuk menyimpan instance database di dalam konteks Echo
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// Endpoint untuk registrasi dan login pengguna
	e.POST("/users/register", controllers.RegisterUser)
	e.POST("/users/login", controllers.LoginUser)

	// Konfigurasi JWT
	config := jwt.Config{
		SigningKey: []byte("secret"),
	}

	// Endpoint yang dilindungi oleh middleware JWT
	// e.GET("/rentals", controllers.GetAllRentals, jwt.WithConfig(config))
	// e.GET("/rentals/active", controllers.GetActiveRentals, jwt.WithConfig(config))
	// e.GET("/vehicles/availability", controllers.GetVehicleAvailability, jwt.WithConfig(config))
}

// SetupRoutes configures the API routes
// func SetupRoutes(router *gin.Engine) {
// 	// Initialize database connection
// 	err := utils.InitializeDatabase()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Define API routes
// 	api := router.Group("/api")

// 	// User routes
// 	api.POST("/users", handlers.CreateUserHandler)
// 	api.GET("/users/:id", handlers.GetUserHandler)

// 	// Game routes
// 	api.POST("/games", handlers.CreateGameHandler)
// 	api.GET("/games/:id", handlers.GetGameHandler)

// 	// Rental routes
// 	api.POST("/rentals", handlers.CreateRentalHandler)
// 	api.GET("/rentals/:id", handlers.GetRentalHandler)

	// Apply authentication middleware to protected routes
	// Example:
	// protected := api.Group("/protected")
	// protected.Use(middleware.AuthMiddleware)
	// protected.GET("/some-route", someHandler)
// }
