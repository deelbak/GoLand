package main

import (
	"log"

	"github.com/deelbak/Goland/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	connectionString := "host=localhost port=5432 dbname=postgres user=postgres password=poiu0987uiop sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{}, &models.Item{})

}
