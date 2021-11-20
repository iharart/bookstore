package handler

func (a *App) SetUpRouters() {
	a.Post("/books", a.CreateBook)
	a.Get("/books/{title}", a.GetBookById)
	a.Get("/books", a.GetAllBooks)
	a.Delete("/books/{title}", a.DeleteBook)
	a.Put("/books/{title}", a.UpdateBook)
}
