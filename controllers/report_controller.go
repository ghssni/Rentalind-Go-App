package controllers

import (
	"net/http"
	"Rentalind-Go-App/handlers"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
)

// BookingReportController handles booking report requests
type BookingReportController struct{}

// GetBookingReport handles getting the booking report
func (c *BookingReportController) GetBookingReport(e echo.Context) error {
	return handlers.GetBookingReport(e)
}

// GetBookingReportUser handles getting the booking report for a specific user
func (c *BookingReportController) GetBookingReportUser(e echo.Context) error {
	return handlers.GetBookingReportUser(e)
}