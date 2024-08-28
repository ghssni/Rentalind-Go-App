package models

// Game struct represents a game in the database
type Game struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Availability bool    `json:"availability"`
	RentalCost  float64 `json:"rental_cost"`
	Category     string  `json:"category"`
}
