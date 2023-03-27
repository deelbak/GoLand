package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/deelbak/go-project-kbtu/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// insert user, items
// get by name, price, rating

func (m *postgresDBRepo) InsertUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	query := `insert into users (first_name, last_name, email, password, role, created_at, updated_at) 
				values ($1, $2, $3, $4, $5, $6, $7) returning id`

	password, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = password

	err = m.DB.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, err
}

func (m *postgresDBRepo) InsertItem(item models.Item) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	query := `insert into item (name, price, rating, number_of_ratings, category, seller_id, created_at, updated_at)
				values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	err := m.DB.QueryRowContext(ctx, query,
		item.Name,
		item.Price,
		0,
		0,
		item.Category,
		item.SellerID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) UpdateItemRating(id int, rating int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update item set rating = rating + (($1 - rating) / (number_of_ratings + 1)),
			number_of_ratings = number_of_ratings + 1,
			updated_at = $2
			where id = $3`

	_, err := m.DB.ExecContext(ctx, query, rating, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) GetAllItems() ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var items []models.Item

	query := `select * from item`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i models.Item
		err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Rating,
			&i.NumberOfRatings,
			&i.Category,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return items, err
		}

		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (m *postgresDBRepo) GetItemByID(id int) (models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var item models.Item

	query := `select * from item where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&item.ID,
		&item.Name,
		&item.Price,
		&item.Rating,
		&item.NumberOfRatings,
		&item.Category,
		&item.SellerID,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return item, err
	}

	return item, nil
}

func (m *postgresDBRepo) GetItemsByName(name string) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var items []models.Item

	query := `select * from item where name = $1`

	rows, err := m.DB.QueryContext(ctx, query, name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i models.Item
		err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Rating,
			&i.NumberOfRatings,
			&i.Category,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return items, err
		}

		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (m *postgresDBRepo) GetItemsByPrice(price float64) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var items []models.Item

	query := `select * from item where price <= $1 order by price asc`

	rows, err := m.DB.QueryContext(ctx, query, price)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i models.Item
		err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Rating,
			&i.NumberOfRatings,
			&i.Category,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return items, err
		}

		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (m *postgresDBRepo) GetItemsByRating(rating float64) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var items []models.Item

	query := `select * from item where rating >= $1 order by rating desc`

	rows, err := m.DB.QueryContext(ctx, query, rating)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i models.Item
		err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Rating,
			&i.NumberOfRatings,
			&i.Category,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return items, err
		}

		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (m *postgresDBRepo) Authenticate(email string, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var gmail string
	var hashedPassword string

	query := `select id, email, password from users where email = $1`
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&id, &gmail, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		log.Println(len([]byte(hashedPassword)))
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
