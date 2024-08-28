package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ghssni/Rentalind-Go-App/models"
	"github.com/ghssni/Rentalind-Go-App/utils"
)

// CreateGameHandler handles the creation of a new game
func CreateGameHandler(c *gin.Context) {
	// Bind JSON data to Game struct
	var game models.Game
	if err := c.BindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new game in the database
	err := utils.CreateGame(&game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Game created successfully", "id": game.ID})
}

// GetGameHandler handles the retrieval of a game by ID
func GetGameHandler(c *gin.Context) {
	// Get game ID from URL parameter
	gameID := c.Param("id")

	// Retrieve game from the database
	game, err := utils.GetGame(gameID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	// Return game details
	c.JSON(http.StatusOK, game)
}
