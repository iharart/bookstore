package router

import (
	"github.com/gorilla/mux"
	database "github.com/iharart/bookstore/app/database"
	handler "github.com/iharart/bookstore/app/handler"
	"net/http"
)

type Provider struct {
	Router *mux.Router
}

func SetUp() *mux.Router {
	w := &Provider{Router: mux.NewRouter()}

	api := &handler.APIEnv{
		DB: database.GetDB(),
	}
	w.Post("/books", api.CreateBook)
	w.Get("/books/{id}", api.GetBookById)
	w.Get("/books", api.GetAllBooks)
	w.Delete("/books/{id}", api.DeleteBook)
	w.Put("/books/{id}", api.UpdateBook)
	return w.Router
}

func (p *Provider) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	p.Router.HandleFunc(path, f).Methods("GET")
}

func (p *Provider) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	p.Router.HandleFunc(path, f).Methods("POST")
}

func (p *Provider) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	p.Router.HandleFunc(path, f).Methods("PUT")
}

func (p *Provider) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	p.Router.HandleFunc(path, f).Methods("DELETE")
}
