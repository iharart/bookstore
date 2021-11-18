package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	book  string = "Book"
	genre string = "Genre"
)

type Tabler interface {
	TableName() string
}

func (Book) TableName() string {
	return book
}

func (Genre) TableName() string {
	return genre
}

type Book struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `json:"name"`
	GenreID int     `json:"-"`
	Price   float64 `json:"price"`
	Amount  uint    `json:"amount"`
	Genre   Genre   `gorm:"foreignKey:GenreID"`
}

type Genre struct {
	ID   int    `gorm:"gorm:primaryKey" json:"id"`
	Name string `json:"name"`
}

func Initialize() {

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

	Connector, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start working with GORM..")

	Connector.AutoMigrate(&Book{}, &Genre{})

	//Get all books
	var books []Book
	fmt.Println("Retrieving all records from Books:")
	if err := Connector.Debug().Preload(genre).Find(&books).Error; err != nil {
		log.Fatal(err)
	}

	result, err := json.Marshal(books)
	fmt.Println(string(result))
}
func main() {
	Initialize()
}
