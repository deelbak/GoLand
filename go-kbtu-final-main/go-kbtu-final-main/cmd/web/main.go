package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"github.com/deelbak/final-go-kbtu/internal/config"
	"github.com/deelbak/final-go-kbtu/internal/handlers"
	"github.com/deelbak/final-go-kbtu/internal/models"
	"github.com/deelbak/final-go-kbtu/internal/render"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const portNumber = ":8081"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("../../static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// read flags
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "postgres", "Database name")
	dbUser := flag.String("dbuser", "postgres", "Database user")
	dbPass := flag.String("dbpass", "poiu0987uiop", "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

	//change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Users{}, &models.Items{}, &models.Comments{}, &models.Roles{}, &models.Categories{})

	//change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	flag.Parse()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Println("zhumys istep zhatyr")
	err = srv.ListenAndServe()
	log.Fatal(err)

}
