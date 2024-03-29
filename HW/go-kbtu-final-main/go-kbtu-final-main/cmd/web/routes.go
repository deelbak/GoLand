package main

import (
	"net/http"

	"github.com/deelbak/final-go-kbtu/internal/config"
	"github.com/deelbak/final-go-kbtu/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/signup", handlers.Repo.ShowSignUp)
	mux.Post("/signup", handlers.Repo.SignUp)

	mux.Get("/login", handlers.Repo.ShowLogin)
	mux.Post("/login", handlers.Repo.Login)
	mux.Get("/logout", handlers.Repo.Logout)

	mux.Get("/items/filter", handlers.Repo.GetAllItems)
	mux.Post("/items/filter", handlers.Repo.SortItems)

	mux.Get("/items/filter/insert", handlers.Repo.ShowInsertItem)
	mux.Post("/items/filter/insert", handlers.Repo.InsertItem)

	mux.Get("/items/filter/update/{id}", handlers.Repo.ShowUpdateItem)
	mux.Post("/items/filter/update/{id}", handlers.Repo.UpdateItem)

	mux.Get("/items/filter/{id}", handlers.Repo.SingleItem)
	mux.Post("/items/filter/{id}", handlers.Repo.PostSingleItem)

	mux.Get("/items/filter/delete/{id}", handlers.Repo.DeleteSingleItem)

	mux.Get("/items/filter/buy/{id}", handlers.Repo.BuyItem)
	mux.Post("/items/filter/buy/{id}", handlers.Repo.PostBuyItem)

	return mux
}
