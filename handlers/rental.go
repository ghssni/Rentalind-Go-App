package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ghssni/Rentalind-Go-App/models"
	"github.com/ghssni/Rentalind-Go-App/utils"
)

// CreateRentalHandler handles the creation of a new rental
func CreateRentalHandler(c *gin.Context) {
	// Bind JSON data to Rental struct
	var rental models.Rental
	if err := c.BindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new rental in the database
	err := utils.CreateRental(&rental)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Rental created successfully", "id": rental.ID})
}

// GetRentalHandler handles the retrieval of a rental by ID
func GetRentalHandler(c *gin.Context) {
	// Get rental ID from URL parameter
	rentalID := c.Param("id")

	// Retrieve rental from the database
	rental, err := utils.GetRental(rentalID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rental not found"})
		return
	}

	// Return rental details
	c.JSON(http.StatusOK, rental)
}
