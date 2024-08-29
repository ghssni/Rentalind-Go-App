package handlers

import (
	"net/http"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
)

// RentProducts handles rental requests
func RentProducts(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid user ID",
		})
	}

	user := new(models.User)
	result := db.First(&user, userID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	rentalID, err := strconv.Atoi(c.FormValue("rental_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid rental ID",
		})
	}

	rental := new(models.Rental)
	result = db.First(&rental, rentalID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Rental not found",
		})
	}

	if rental.Availability == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Rental is not available",
		})
	}

	rentalHistory := new(models.RentalHistory)
	rentalHistory.UserID = user.ID
	rentalHistory.RentalID = rental.ID
	rentalHistory.PaymentID = 0 // assume payment ID is 0 for now
	rentalHistory.RentalStartDate = time.Now()
	rentalHistory.RentalEndDate = time.Now().AddDate(0, 0, 7) // assume 7-day rental period

	result = db.Create(&rentalHistory)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create rental history",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Rental successful",
		"rental_id": rentalHistory.ID,
	})
}
	
// GetBookingReport handles getting the booking report
func GetBookingReport(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	var rentalHistories []models.RentalHistory
	result := db.Find(&rentalHistories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve rental histories",
		})
	}

	var bookingReport []map[string]interface{}
	for _, rentalHistory := range rentalHistories {
		bookingReportItem := map[string]interface{}{
			"rental_id": rentalHistory.RentalID,
			"user_id":   rentalHistory.UserID,
			"start_date": rentalHistory.RentalStartDate,
			"end_date":   rentalHistory.RentalEndDate,
		}
		bookingReport = append(bookingReport, bookingReportItem)
	}

return c.JSON(http.StatusOK, echo.Map{
	"booking_report": bookingReport,
})
}
	
// GetBookingReportUser handles getting the booking report for a specific user
func GetBookingReportUser(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid user ID",
		})
	}

	var rentalHistories []models.RentalHistory
	result := db.Where("user_id = ?", userID).Find(&rentalHistories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve rental histories",
		})
	}

	var bookingReport []map[string]interface{}
	for _, rentalHistory := range rentalHistories {
		bookingReportItem := map[string]interface{}{
			"rental_id": rentalHistory.RentalID,
			"start_date": rentalHistory.RentalStartDate,
			"end_date":   rentalHistory.RentalEndDate,
		}
		bookingReport = append(bookingReport, bookingReportItem)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"booking_report": bookingReport,
	})
}