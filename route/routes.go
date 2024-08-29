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
	e.POST("/topup/:user_id", handlers.TopUp)
	e.POST("/topup/user/:user_id", handlers.TopupUser)
	e.POST("/rentals/:user_id", handlers.RentProducts)
	e.GET("/booking-report", handlers.GetBookingReport)
	e.GET("/booking-report/user/:user_id", handlers.GetBookingReportUser)
}
