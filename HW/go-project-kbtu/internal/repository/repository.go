package repository

import "github.com/deelbak/go-project-kbtu/internal/models"

type DatabaseRepo interface {
	InsertUser(user models.User) (int, error)
	InsertItem(item models.Item) (int, error)
	UpdateItemRating(id int, rating int) error
	GetAllItems() ([]models.Item, error)
	GetItemByID(id int) (models.Item, error)
	GetItemsByPrice(price float64) ([]models.Item, error)
	GetItemsByName(name string) ([]models.Item, error)
	GetItemsByRating(rating float64) ([]models.Item, error)
	Authenticate(email string, password string) (int, string, error)
}
