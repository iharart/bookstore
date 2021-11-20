package handler

import (
	model "bookstore/model"
	utils "bookstore/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	book := model.Book{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&book).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(false, w, http.StatusCreated, book)
}

func GetBookById(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sId := vars["id"]
	id := utils.StringToUint(sId)
	employee := getEmployeeOr404(db, id, w, r)
	if employee == nil {
		return
	}
	utils.RespondJSON(false, w, http.StatusOK, employee)
}

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	books := []model.Book{}
	if err := db.Debug().Preload(model.GENRE).Order("name").Find(&books).Error; err != nil {
		log.Fatal(err)
	}
	utils.RespondJSON(true, w, http.StatusOK, books)
}

func getEmployeeOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Book {
	book := model.Book{}
	if err := db.First(&book, model.Book{ID: id}).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}
