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
	book := getEntityOr404(db, w, r)
	if book == nil {
		return
	}
	utils.RespondJSON(false, w, http.StatusOK, book)
}

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	books := []model.Book{}
	if err := db.Debug().Preload(model.GENRE).Order("name").Find(&books).Error; err != nil {
		log.Fatal(err)
	}
	utils.RespondJSON(true, w, http.StatusOK, books)
}

func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	book := getEntityOr404(db, w, r)
	if book == nil {
		return
	}
	if err := db.Delete(&book).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(false, w, http.StatusNoContent, nil)
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	book := getEntityOr404(db, w, r)
	if book == nil {
		return
	}

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
	utils.RespondJSON(false, w, http.StatusOK, book)
}

func getEntityOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request) *model.Book {
	book := model.Book{}
	id := getId(r)
	if err := db.First(&book, model.Book{ID: id}).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}

func getId(r *http.Request) uint {
	vars := mux.Vars(r)
	sId := vars["id"]
	return utils.StringToUint(sId)
}
