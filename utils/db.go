package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/ghssni/Rentalind-Go-App/config"
	"github.com/ghssni/Rentalind-Go-App/models"
)

// DB is the global database connection
var DB *sql.DB

// InitializeDatabase connects to the database
func InitializeDatabase() error {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Create a database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Test database connection
	err = db.Ping()
	if err != nil {
		return err
	}

	// Set global database connection
	DB = db
	return nil
}

// CreateUser creates a new user in the database
func CreateUser(user *models.User) error {
	// Insert user into database
	_, err := DB.Exec("INSERT INTO users (email, password, deposit_amount) VALUES (?, ?, ?)", user.Email, user.Password, user.DepositAmount)
	if err != nil {
		return err
	}

	// Get the last inserted ID
	userID, err := DB.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(userID)
	return nil
}

// GetUser retrieves a user from the database by ID
func GetUser(userID string) (*models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT id, email, password, deposit_amount FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Email, &user.Password, &user.DepositAmount)
	if err != nil {
		return nil, err
	}

	// Retrieve rental history for user
	err = getUserRentalHistory(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateGame creates a new game in the database
func CreateGame(game *models.Game) error {
	// Insert game into database
	_, err := DB.Exec("INSERT INTO games (name, availability, rental_cost, category) VALUES (?, ?, ?, ?)", game.Name, game.Availability, game.RentalCost, game.Category)
	if err != nil {
		return err
	}

	// Get the last inserted ID
	gameID, err := DB.LastInsertId()
	if err != nil {
		return err
	}

	game.ID = int(gameID)
	return nil
}

// GetGame retrieves a game from the database by ID
func GetGame(gameID string) (*models.Game, error) {
	var game models.Game
	err := DB.QueryRow("SELECT id, name, availability, rental_cost, category FROM games WHERE id = ?", gameID).Scan(&game.ID, &game.Name, &game.Availability, &game.RentalCost, &game.Category)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// CreateRental creates a new rental in the database
func CreateRental(rental *models.Rental) error {
	// Insert rental into database
	_, err := DB.Exec("INSERT INTO rentals (user_id, game_id, rental_date, return_date) VALUES (?, ?, ?, ?)", rental.UserID, rental.GameID, rental.RentalDate, rental.ReturnDate)
	if err != nil {
		return err
	}

	// Get the last inserted ID
	rentalID, err := DB.LastInsertId()
	if err != nil {
		return err
	}

	rental.ID = int(rentalID)
	return nil
}

// GetRental retrieves a rental from the database by ID
func GetRental(rentalID string) (*models.Rental, error) {
	var rental models.Rental
	err := DB.QueryRow("SELECT id, user_id, game_id, rental_date, return_date FROM rentals WHERE id = ?", rentalID).Scan(&rental.ID, &rental.UserID, &rental.GameID, &rental.RentalDate, &rental.ReturnDate)
	if err != nil {
		return nil, err
	}

	return &rental, nil
}

// getUserRentalHistory retrieves the rental history for a given user
func getUserRentalHistory(user *models.User) error {
	rows, err := DB.Query("SELECT id, user_id, game_id, rental_date, return_date FROM rentals WHERE user_id = ?", user.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var rental models.Rental
		err := rows.Scan(&rental.ID, &rental.UserID, &rental.GameID, &rental.RentalDate, &rental.ReturnDate)
		if err != nil {
			return err
		}
		user.RentalHistory = append(user.RentalHistory, rental)
	}

	return nil
}
