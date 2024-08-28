package models

// User struct represents a user in the database
type User struct {
	ID             int     `json:"id"`
	Name		   string  `json:"name"`
	address		   string  `json:"address"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	DepositAmount  float64 `json:"deposit_amount"`
	RentalHistory []Rental `json:"rental_history"` // Add rental history for user
}
