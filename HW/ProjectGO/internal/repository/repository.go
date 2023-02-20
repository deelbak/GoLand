package repository

import "github.com/nanmenkaimak/bookstore/internal/models"

type DatabaseRepo interface {
	InsertBook(book models.Book) (int, error)
	GetAllBooks() ([]models.Book, error)
	GetSoldBooks() ([]models.Book, error)
	GetNotSoldBooks() ([]models.Book, error)
	GetBooksByAuthor(author string) ([]models.Book, error)
	GetBooksByPrice(price int) ([]models.Book, error)
	GetBooksByName(bookName string) ([]models.Book, error)
	UpdateBookSoldStatus(id int) error
	UpdateBookPrice(id int, price int) error
	DeleteBook(id int) error

	InsertUser(user models.User) (int, error)
	DeleteUser(id int) error
	UpdateUserBalance(id int, profit_money int) error
	UpdateUser(user models.User) error
}