package models

// Book struct represents a Book in the database
type Book struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Availability bool    `json:"availability"`
	RentalCost  float64 `json:"rental_cost"`
	Category     string  `json:"category"`
}

