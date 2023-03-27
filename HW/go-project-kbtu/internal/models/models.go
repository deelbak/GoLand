package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Item struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Price           float64   `json:"price"`
	Rating          float64   `json:"rating"`
	NumberOfRatings int       `json:"number_of_ratings"`
	Category        string    `json:"category"`
	SellerID        int       `json:"sellerid"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
