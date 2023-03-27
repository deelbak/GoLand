package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/deelbak/go-project-kbtu/internal/config"
	"github.com/deelbak/go-project-kbtu/internal/driver"
	"github.com/deelbak/go-project-kbtu/internal/forms"
	"github.com/deelbak/go-project-kbtu/internal/models"
	"github.com/deelbak/go-project-kbtu/internal/render"
	"github.com/deelbak/go-project-kbtu/internal/repository"
	"github.com/deelbak/go-project-kbtu/internal/repository/dbrepo"
)

var Repo *Repository // repository used by handlers

type Repository struct { // repository type
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository { //creates new repository
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.gohtml", &models.TemplateData{
		Form: nil,
	})
}

func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		render.Template(w, r, "login.page.gohtml", &models.TemplateData{
			Form: form,
		})
		return
	}
	id, _, err := m.DB.Authenticate(email, password)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	fmt.Println("login isted")

	m.App.Session.Put(r.Context(), "user_id", id)
	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}

func (m *Repository) ShowSignUp(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "signup.page.gohtml", &models.TemplateData{
		Form: nil,
	})
}

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	first_name := r.Form.Get("firstname")
	last_name := r.Form.Get("lastname")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	role, err := strconv.Atoi(r.Form.Get("role"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	newUser := models.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
		Role:      role,
	}
	form := forms.New(r.PostForm)

	form.IsEmail("email")
	form.MinLength("password", 8)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["user"] = newUser
		render.Template(w, r, "signup.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	_, err = m.DB.InsertUser(newUser)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("user logged")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (m *Repository) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := m.DB.GetAllItems()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["items"] = items
	render.Template(w, r, "filter.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

// SortItems filters by name, rating, price
func (m *Repository) SortItems(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var items []models.Item

	if r.Form.Get("price") != "" {
		price, err := strconv.ParseFloat(r.Form.Get("price"), 64)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		items, err = m.DB.GetItemsByPrice(price)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Form.Get("rating") != "" {
		rating, err := strconv.ParseFloat(r.Form.Get("rating"), 64)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		items, err = m.DB.GetItemsByRating(rating)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Form.Get("name") != "" {
		name := r.Form.Get("name")

		items, err = m.DB.GetItemsByName(name)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	data := make(map[string]interface{})
	data["items"] = items

	render.Template(w, r, "filter.page.gohtml", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) SingleItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	item, err := m.DB.GetItemByID(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["item"] = item

	render.Template(w, r, "item.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

func (m *Repository) PostSingleItem(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	rate, err := strconv.Atoi(r.Form.Get("rate"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = m.DB.UpdateItemRating(id, rate)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}
