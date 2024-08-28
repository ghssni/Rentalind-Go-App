package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ghssni/Rentalind-Go-App/models"
	"github.com/ghssni/Rentalind-Go-App/utils"
)

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(c *gin.Context) {
	// Bind JSON data to User struct
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user in the database
	err := utils.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "id": user.ID})
}

// GetUserHandler handles the retrieval of a user by ID
func GetUserHandler(c *gin.Context) {
	// Get user ID from URL parameter
	userID := c.Param("id")

	// Retrieve user from the database
	user, err := utils.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user details
	c.JSON(http.StatusOK, user)
}
