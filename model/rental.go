package models

// Rental struct represents a rental in the database
type Rental struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	GameID      int    `json:"game_id"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
}
