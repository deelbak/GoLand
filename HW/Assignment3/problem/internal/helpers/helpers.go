// package helpers

// import (
// 	"net/http"

// 	"github.com/deelbak/go-project-kbtu/internal/config"
// )

// var app *config.AppConfig

// // NewHelpers sets up app config for helpers
// func NewHelpers(a *config.AppConfig) {
// 	app = a
// }

// func IsAuthenticated(r *http.Request) bool {
// 	exists := app.Session.Exists(r.Context(), "user_id")
// 	return exists
// }