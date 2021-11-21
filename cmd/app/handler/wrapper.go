package handler

import (
	"log"
	"net/http"
)

func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook(a.DB, w, r)
}

func (a *App) GetBookById(w http.ResponseWriter, r *http.Request) {
	GetBookById(a.DB, w, r)
}

func (a *App) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	GetAllBooks(a.DB, w, r)
}

func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook(a.DB, w, r)
}

func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	DeleteBook(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
