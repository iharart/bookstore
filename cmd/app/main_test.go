package main

import (
	"bookstore/handler"
	"bookstore/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetBooks_EmptyResult(t *testing.T) {
	service := &handler.Service{}
	service.Initialize()
	req, w := setGetBooksRouter(service.DB)
	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := model.Book{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := model.Book{}
	a.NotEqual(expected, actual)
}

func setGetBooksRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	service := &handler.Service{DB: db}
	service.Router = mux.NewRouter()
	service.Get("/books", service.GetAllBooks)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	service.Router.ServeHTTP(w, req)
	return req, w
}
