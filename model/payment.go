package models

type Payment struct{
	ID             int     `json:"id"`
	Amount  	   float64 `json:"amount"`
	Status 		   string  `json:"status"`
	Url 		   string  `json:"url"`
}