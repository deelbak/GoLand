package models

type User struct {
	ID       int    `gorm:"primaryKey; autpIncrement 1"`
	Name     string `gorm:"not null"`
	Surname  string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Role     string `gorm:"not null"`
}
type Item struct {
	ID              int     `gorm:"primaryKey; autpIncrement 1"`
	Name            string  `gorm:"not null"`
	Price           float64 `gorm:"not null"`
	Rating          float64 `gorm:"not null"`
	Description     string  `gorm:"not null"`
	NumberOfRatings int     `gorm:"not null"`
	SellerID        int     `gorm:"not null"`
	Seller          User    `gorm:"not null"`
}
