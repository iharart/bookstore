package handler

import (
	model "bookstore/model"
	utils "bookstore/utils"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var books []model.Book
	fmt.Println("Retrieving all records from Books:")
	if err := db.Debug().Preload(model.GENRE).Find(&books).Error; err != nil {
		log.Fatal(err)
	}
	utils.RespondJSON(w, http.StatusOK, books)
}
