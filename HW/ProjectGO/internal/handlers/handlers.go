package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nanmenkaimak/bookstore/internal/driver"
	"github.com/nanmenkaimak/bookstore/internal/models"
	"github.com/nanmenkaimak/bookstore/internal/repository"
	"github.com/nanmenkaimak/bookstore/internal/repository/dbrepo"
	"mime"
	"net/http"
	"strconv"
)

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

func (m *Repository) InsertBook(w http.ResponseWriter, r *http.Request) {
	// Enforce a JSON Content-Type.
	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "expected application/json Content-type", http.StatusUnsupportedMediaType)
		return
	}

	var books models.Book

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&books); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = m.DB.InsertBook(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := m.DB.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) GetSoldBooks(w http.ResponseWriter, r *http.Request) {
	books, err := m.DB.GetSoldBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) GetNotSoldBooks(w http.ResponseWriter, r *http.Request) {
	books, err := m.DB.GetNotSoldBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	books, err := m.DB.GetBooksByAuthor(mux.Vars(r)["author"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) GetBooksByPrice(w http.ResponseWriter, r *http.Request) {
	price, err := strconv.Atoi(mux.Vars(r)["price"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books, err := m.DB.GetBooksByPrice(price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) GetBooksByName(w http.ResponseWriter, r *http.Request) {
	books, err := m.DB.GetBooksByName(mux.Vars(r)["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, books)
}

func (m *Repository) UpdateBookSoldStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = m.DB.UpdateBookSoldStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (m *Repository) UpdateBookPrice(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("osy zher")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var books models.Book

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&books); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = m.DB.UpdateBookPrice(id, books.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (m *Repository) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = m.DB.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}