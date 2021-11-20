package handler

import (
	"bookstore/model"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	connectionString := "admin:admin@tcp(db:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Retrieve only one row by sql")
	id := 1
	var name string
	if err := sqlDb.QueryRow("SELECT name FROM Genre WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start working with GORM..")
	a.DB = model.Migrate(db)
	a.Router = mux.NewRouter()
	a.SetUpRouters()

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
