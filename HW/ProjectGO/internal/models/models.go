package models

import "time"

type Book struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	IsSold    bool      `json:"issold"`
	SellerID  int       `json:"sellerid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type User struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AccessLevel int       `json:"accesslevel"`
	Balance     int       `json:"balance"`
	BoughtBooks []Book    `json:"boughtbooks"`
	SellBooks   []Book    `json:"sellbooks"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}