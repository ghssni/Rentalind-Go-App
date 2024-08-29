package controllers

import (
	"net/http"
	"Rentalind-Go-App/handlers"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
)

// TopUpController handles top-up requests
type TopUpController struct{}

// TopUp handles top-up requests
func (c *TopUpController) TopUp(e echo.Context) error {
	return handlers.TopUp(e)
}

// TopupUser handles top-up requests for a specific user
func (c *TopUpController) TopupUser(e echo.Context) error {
	return handlers.TopupUser(e)
}