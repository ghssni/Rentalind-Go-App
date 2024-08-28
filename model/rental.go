package models

// Rental struct represents a rental in the database
type Rental struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	BookID      int    `json:"book_id"`
	PaymentID	int	   `json:"payment_id"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
}
