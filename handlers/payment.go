package handlers

import (
	"net/http"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
)

// TopUp handles top-up requests
func TopUp(c echo.Context) error {
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

	depositAmount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid deposit amount",
		})
	}

	user.Deposit += depositAmount
	result = db.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update user deposit",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Top-up successful",
		"deposit": user.Deposit,
	})
}

// TopupUser handles top-up requests for a specific user
func TopupUser(c echo.Context) error {
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

	depositAmount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid deposit amount",
		})
	}

	user.Deposit += depositAmount
	result = db.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update user deposit",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Top-up successful",
		"deposit": user.Deposit,
	})
}