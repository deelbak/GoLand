package dbrepo

import (
	"context"
	"github.com/nanmenkaimak/bookstore/internal/models"
	"time"
)

// ID, name, price, sold/not sold, date, author
// books insert, delete, update, get by author, get by price, get by name
// user insert, delete, update balance, accesslevel, password, get books, get sold/not sold books

// InsertBook insert a book to the bookstore database
func (m *postgresDBRepo) InsertBook(book models.Book) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	query := `insert into bookstore (author, name, price, issold, sellerid, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err := m.DB.QueryRowContext(ctx, query,
		book.Author,
		book.Name,
		book.Price,
		book.IsSold,
		book.SellerID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// GetAllBooks gets all books from the database
func (m *postgresDBRepo) GetAllBooks() ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var books []models.Book

	query := `select * from bookstore order by id`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Book
		err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Name,
			&i.Price,
			&i.IsSold,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return books, err
		}

		books = append(books, i)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

// GetSoldBooks gets sold books from the database
func (m *postgresDBRepo) GetSoldBooks() ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var books []models.Book

	query := `select * from bookstore where issold = true order by id`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Book
		err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Name,
			&i.Price,
			&i.IsSold,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return books, err
		}

		books = append(books, i)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

// GetNotSoldBooks gets not sold books from the database
func (m *postgresDBRepo) GetNotSoldBooks() ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var books []models.Book

	query := `select * from bookstore where issold = false order by id`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Book
		err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Name,
			&i.Price,
			&i.IsSold,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return books, err
		}

		books = append(books, i)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

// GetBooksByAuthor gets all books by author
func (m *postgresDBRepo) GetBooksByAuthor(author string) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var books []models.Book

	query := `select * from bookstore where author = $1`

	rows, err := m.DB.QueryContext(ctx, query, author)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Book
		err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Name,
			&i.Price,
			&i.IsSold,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return books, err
		}

		books = append(books, i)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

// GetBooksByPrice gets all books by price
func (m *postgresDBRepo) GetBooksByPrice(price int) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var books []models.Book

	query := `select * from bookstore where price <= $1`

	rows, err := m.DB.QueryContext(ctx, query, price)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Book
		err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Name,
			&i.Price,
			&i.IsSold,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return books, err
		}
		books = append(books, i)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

// GetBooksByName gets all books by name
func (m *postgresDBRepo) GetBooksByName(bookName string) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var books []models.Book

	query := `select * from bookstore where name = $1`

	rows, err := m.DB.QueryContext(ctx, query, bookName)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Book
		err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Name,
			&i.Price,
			&i.IsSold,
			&i.SellerID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return books, err
		}
		books = append(books, i)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

// UpdateBookSoldStatus updates selling status of the book
func (m *postgresDBRepo) UpdateBookSoldStatus(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update bookstore set issold = true, updated_at = $1 where id = $2`

	_, err := m.DB.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBookPrice updates price of the book
func (m *postgresDBRepo) UpdateBookPrice(id int, price int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update bookstore set price = $1, updated_at = $2 where id = $3`

	_, err := m.DB.ExecContext(ctx, query, price, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook deletes book for database
func (m *postgresDBRepo) DeleteBook(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from bookstore where id = $1`
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// InsertUser insert a new user to the bookstore database
func (m *postgresDBRepo) InsertUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	query := `insert into users (first_name, last_name, email, password, accesslevel, balance, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	err := m.DB.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.AccessLevel,
		user.Balance,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// DeleteUser deletes user from database
func (m *postgresDBRepo) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from users where id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserBalance updates balance of the user
func (m *postgresDBRepo) UpdateUserBalance(id int, profit_money int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update users set balance = balance + $1, updated_at = $2 where id = $3`

	_, err := m.DB.ExecContext(ctx, query, profit_money, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates a user in the database
func (m *postgresDBRepo) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update users set first_name = $1, last_name = $2, email = $3, password = $4, balance = $5, updated_at = $6 where id = $7`

	_, err := m.DB.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Balance,
		time.Now(),
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}