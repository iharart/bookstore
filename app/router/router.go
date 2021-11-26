package router

import (
	"github.com/gorilla/mux"
	database "github.com/iharart/bookstore/app/database"
	handler "github.com/iharart/bookstore/app/handler"
	"net/http"
)

type Wrapper struct {
	Router *mux.Router
}

func SetUp() *mux.Router {
	w := &Wrapper{Router: mux.NewRouter()}

	api := &handler.APIEnv{
		DB: database.GetDB(),
	}
	w.Post("/book", api.CreateBook)
	w.Get("/books/{title}", api.GetBookById)
	w.Get("/books", api.GetAllBooks)
	w.Delete("/books/{title}", api.DeleteBook)
	w.Put("/books/{title}", api.UpdateBook)
	return w.Router
}

func (w *Wrapper) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	w.Router.HandleFunc(path, f).Methods("GET")
}

func (w *Wrapper) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	w.Router.HandleFunc(path, f).Methods("POST")
}

func (w *Wrapper) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	w.Router.HandleFunc(path, f).Methods("PUT")
}

func (w *Wrapper) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	w.Router.HandleFunc(path, f).Methods("DELETE")
}
