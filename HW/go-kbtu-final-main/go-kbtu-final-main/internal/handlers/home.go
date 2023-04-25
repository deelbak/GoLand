package handlers

import (
	"net/http"

	"github.com/deelbak/final-go-kbtu/internal/models"
	"github.com/deelbak/final-go-kbtu/internal/render"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.gohtml", &models.TemplateData{})
}
