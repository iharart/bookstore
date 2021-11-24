package handler

import (
	"log"
	"net/http"
)

func (s *Service) CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook(s.DB, w, r)
}

func (s *Service) GetBookById(w http.ResponseWriter, r *http.Request) {
	GetBookById(s.DB, w, r)
}

func (s *Service) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	GetAllBooks(s.DB, w, r)
}

func (s *Service) UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook(s.DB, w, r)
}

func (s *Service) DeleteBook(w http.ResponseWriter, r *http.Request) {
	DeleteBook(s.DB, w, r)
}

func (s *Service) Run(host string) {
	log.Fatal(http.ListenAndServe(host, s.Router))
}
