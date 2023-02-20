package dbrepo

import (
	"database/sql"
	"github.com/nanmenkaimak/bookstore/internal/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}