package main

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Item struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Rating      float64
}

type System struct {
	db *sql.DB
}

func NewSystem(db *sql.DB) *System {
	return &System{db: db}
}

func (s *System) RegisterUser(name string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("INSERT INTO usersofmagaz (name, email, password_hash) VALUES ($1, $2, $3)", name, email, hashedPassword)
	return err
}

func (s *System) AuthenticateUser(email string, password string) (*User, error) {
	var user User
	err := s.db.QueryRow("SELECT id, name, email, password_hash FROM usersofmagaz WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *System) SearchItemsByName(name string) ([]Item, error) {
	rows, err := s.db.Query("SELECT id, name, description, price, rating FROM items WHERE LOWER(name) LIKE LOWER($1)", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Rating)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *System) FilterItemsByPriceAndRating(minPrice float64, maxPrice float64, minRating float64, maxRating float64) ([]Item, error) {
	rows, err := s.db.Query("SELECT id, name, description, price, rating FROM items WHERE price >= $1 AND price <= $2 AND rating >= $3 AND rating <= $4", minPrice, maxPrice, minRating, maxRating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Rating)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *System) GiveItemRating(itemID int, userID int, rating float64) error {
	_, err := s.db.Exec("INSERT INTO items (item_id, user_id, rating) VALUES ($1, $2, $3) ON CONFLICT (item_id, user_id) DO UPDATE SET rating = (SELECT AVG(rating) FROM item_ratings WHERE item_id = $1) WHERE id = $1", itemID)
return err
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:poiu0987uiop@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	system := NewSystem(db)

	// err = system.RegisterUser("Adilbek Zhumazhanov", "a_zhumazhanov@example.com", "qwerty")
	// if err != nil {
	// 	panic(err)
	// }

	user, err := system.AuthenticateUser("a_zhumazhanov@example.com", "qwerty")
	if err != nil {
		panic(err)
	}

	items, err := system.SearchItemsByName("phone")
	if err != nil {
		panic(err)
	}
	fmt.Println(items)

	items, err = system.FilterItemsByPriceAndRating(100, 1000, 4, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(items)

	err = system.GiveItemRating(1, user.ID, 4.7)
	if err != nil {
		panic(err)
	}
}