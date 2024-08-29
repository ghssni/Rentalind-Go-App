package controllers

import (
	"net/http"
	"Rentalind-Go-App/handlers"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
)

// RentalController handles rental requests
type RentalController struct{}

// RentProducts handles rental requests
func (c *RentalController) RentProducts(e echo.Context) error {
	return handlers.RentProducts(e)
}