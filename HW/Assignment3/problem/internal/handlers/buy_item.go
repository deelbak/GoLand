package handlers

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/deelbak/go-project-kbtu/internal/models"
	"github.com/deelbak/go-project-kbtu/internal/render"
)

func (m *Repository) BuyItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	item, err := m.DB.GetItemById(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["item"] = item

	render.Template(w, r, "buy.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

func (m *Repository) PostBuyItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = m.DB.DeleteItem(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}
