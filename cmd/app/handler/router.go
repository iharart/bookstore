package handler

import "net/http"

func (a *App) SetUpRouters() {
	a.Post("/books", a.CreateBook)
	a.Get("/books/{title}", a.GetBookById)
	a.Get("/books", a.GetAllBooks)
	a.Delete("/books/{title}", a.DeleteBook)
	a.Put("/books/{title}", a.UpdateBook)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
