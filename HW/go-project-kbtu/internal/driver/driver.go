package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	// Set the maximum idle timeout.
	db.SetConnMaxLifetime(maxDbLifetime)
	// Set the maximum number of open (in-use + idle) connections in the pool.
	db.SetMaxOpenConns(maxOpenDbConn)
	// Set the maximum number of idle connections in the pool.
	db.SetMaxIdleConns(maxIdleDbConn)

	dbConn.SQL = db

	err = TestDB(db)
	if err != nil {
		return nil, err
	}

	return dbConn, err
}

// TestDB tries to ping the database
func TestDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
