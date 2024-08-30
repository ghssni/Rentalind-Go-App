package routes

import (
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
	topUpController := &controllers.TopUpController{}
	rentalController := &controllers.RentalController{}
	bookingReportController := &controllers.BookingReportController{}

	e.POST("/topup/:user_id", topUpController.TopUp, jwt.WithConfig(config))
	e.POST("/topup/user/:user_id", topUpController.TopupUser, jwt.WithConfig(config))
	e.POST("/rentals/:user_id", rentalController.RentProducts, jwt.WithConfig(config))
	e.GET("/booking-report", bookingReportController.GetBookingReport, jwt.WithConfig(config))
	e.GET("/booking-report/user/:user_id", bookingReportController.GetBookingReportUser, jwt.WithConfig(config))
	// Book routes
	bookRouter := e.Group("/books")
	bookRouter.POST("", handlers.CreateBook, jwt.WithConfig(config))
	bookRouter.GET("/:book_id", handlers.GetBook, jwt.WithConfig(config))
	bookRouter.GET("", handlers.GetAllBooks, jwt.WithConfig(config))
}
