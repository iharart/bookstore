package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	database "github.com/iharart/bookstore/app/database"
	"github.com/iharart/bookstore/app/model"
	"github.com/iharart/bookstore/app/utils"
	"gorm.io/gorm"
	"net/http"
)

const (
	RecordNotFound      = "record not found"
	BadRequest          = "bad request"
	InternalServerError = "internal server error"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	book, exists, err := database.GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !exists {
		utils.RespondError(w, http.StatusNotFound, RecordNotFound)
		return
	}

	utils.RespondJSON(true, w, http.StatusOK, book)
}

func (a *APIEnv) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.GetAllBooks(a.DB)
	if err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	utils.RespondJSON(true, w, http.StatusOK, books)
}

func (a *APIEnv) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	_, exists, err := database.GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !exists {
		utils.RespondError(w, http.StatusNotFound, RecordNotFound)
		return
	}

	updatedBook := model.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedBook); err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusBadRequest, BadRequest) // or http.StatusInternalServerError?
		return
	}
	defer r.Body.Close()

	if err := database.UpdateBook(a.DB, &updatedBook); err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	a.GetBookById(w, r)
}

func (a *APIEnv) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	_, exists, err := database.GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !exists {
		utils.RespondError(w, http.StatusNotFound, RecordNotFound)
		return
	}

	err = database.DeleteBook(id, a.DB)
	if err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	utils.RespondJSON(true, w, http.StatusOK, nil)
}

func (a *APIEnv) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := model.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		fmt.Println(err)
		utils.RespondError(w, http.StatusBadRequest, BadRequest)
	}
	if err := a.DB.Create(&book).Error; err != nil {
		utils.RespondError(w, http.StatusBadRequest, BadRequest)
		return
	}
	utils.RespondJSON(true, w, http.StatusOK, book.ID)
}

func getId(r *http.Request) uint {
	vars := mux.Vars(r)
	sId := vars["id"]
	return utils.StringToUint(sId)
}
