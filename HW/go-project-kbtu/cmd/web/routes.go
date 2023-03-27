package main

import (
	"net/http"

	"github.com/deelbak/go-project-kbtu/internal/config"
	"github.com/deelbak/go-project-kbtu/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostLogin)

	mux.Get("/user/signup", handlers.Repo.ShowSignUp)
	mux.Post("/user/signup", handlers.Repo.SignUp)

	mux.Get("/items/filter", handlers.Repo.GetAllItems)
	mux.Post("/items/filter", handlers.Repo.SortItems)

	mux.Get("/items/filter/{id}", handlers.Repo.SingleItem)
	mux.Post("/items/filter/{id}", handlers.Repo.PostSingleItem)

	return mux
}
