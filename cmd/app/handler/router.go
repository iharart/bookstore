package handler

import "net/http"

func (s *Service) SetUpRouters() {
	s.Post("/books", s.CreateBook)
	s.Get("/books/{title}", s.GetBookById)
	s.Get("/books", s.GetAllBooks)
	s.Delete("/books/{title}", s.DeleteBook)
	s.Put("/books/{title}", s.UpdateBook)
}

func (s *Service) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("GET")
}

func (s *Service) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("POST")
}

func (s *Service) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("PUT")
}

func (s *Service) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("DELETE")
}
