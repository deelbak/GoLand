package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nanmenkaimak/bookstore/internal/driver"
	"github.com/nanmenkaimak/bookstore/internal/handlers"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"

func main() {
	content, err := os.ReadFile("internal/password.txt")
	if err != nil {
		log.Fatal(err)
	}
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "pp2example", "Database name")
	dbUser := flag.String("dbuser", "pp2example_user", "Database user")
	dbPass := flag.String("dbpass", string(content), "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)

	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer db.SQL.Close()

	flag.Parse()

	router := mux.NewRouter()
	server := handlers.NewRepo(db)

	router.HandleFunc("/books/", server.InsertBook).Methods("POST")
	router.HandleFunc("/books/", server.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{author}/", server.GetBooksByAuthor).Methods("GET")
	router.HandleFunc("/books/price/{price}/", server.GetBooksByPrice).Methods("GET")
	router.HandleFunc("/books/status/{issold}/", server.GetSoldBooks).Methods("GET")
	router.HandleFunc("/books/status/{issold}/", server.GetNotSoldBooks).Methods("GET")
	router.HandleFunc("/books/name/{name}/", server.GetBooksByName).Methods("GET")
	router.HandleFunc("/books/statusu/{id}/", server.UpdateBookSoldStatus).Methods("PUT")
	router.HandleFunc("/books/priceid/{id}/", server.UpdateBookPrice).Methods("PUT")
	router.HandleFunc("/books/delete/{id}/", server.DeleteBook).Methods("DELETE")

	router.StrictSlash(true)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: router,
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}