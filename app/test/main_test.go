package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/iharart/bookstore/app/database"
	_ "github.com/iharart/bookstore/app/database"
	"github.com/iharart/bookstore/app/handler"
	"github.com/iharart/bookstore/app/model"
	"github.com/iharart/bookstore/app/router"
	"github.com/iharart/bookstore/app/utils"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestSuiteEnv struct {
	suite.Suite
	pool     *dockertest.Pool
	resource *dockertest.Resource
	sqlDb    *sql.DB
	api      handler.APIEnv
	provider router.Provider
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteEnv))
}

func (s *TestSuiteEnv) SetupSuite() {
	s.Initialize()
}

func (s *TestSuiteEnv) TearDownSuite() {
	if err := s.sqlDb.Close(); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	if err := s.pool.Purge(s.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *TestSuiteEnv) Initialize() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7",
		Env: []string{
			"MYSQL_ALLOW_EMPTY_PASSWORD=true",
			"MYSQL_ROOT_PASSWORD=admin",
			"MYSQL_DATABASE=bookstore",
			"MYSQL_USER=admin",
			"MYSQL_PASSWORD=admin",
		},
		Tty: true,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		s.sqlDb, err = sql.Open("mysql", fmt.Sprintf("admin:admin@(localhost:%s)/bookstore?charset=utf8mb4&parseTime=True&loc=Local",
			resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return s.sqlDb.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: s.sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.ErrorCheck(err)

	s.api.DB, err = database.Migrate(db)

	utils.ErrorCheck(err)

	s.pool = pool
	s.resource = resource
}

func (s *TestSuiteEnv) TestAddGenresOk() {
	if err := s.api.DB.Save(&SampleGenresOk).Error; err != nil {
		s.Fail("Fail TestAddGenresOk", err)
	}
}

func (s *TestSuiteEnv) TestAddGenresFail() {
	if err := s.api.DB.Create(&SampleGenresFail).Error; err == nil {
		s.Fail("Fail TestAddGenresFail", err)
	}
}

func (s *TestSuiteEnv) TestGetBooksEmptyResult() {
	s.ClearTable(&model.Book{})
	req, w := setGetBooksRouter(s)
	a := s.Assert()

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
	a.Equal(expected, actual)

}

func setGetBooksRouter(s *TestSuiteEnv) (*http.Request, *httptest.ResponseRecorder) {
	s.provider.Router = mux.NewRouter()
	s.provider.Get("/books", s.api.GetAllBooks)
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.provider.Router.ServeHTTP(w, req)
	return req, w
}

func (s *TestSuiteEnv) ClearTable(payload interface{}) {
	s.api.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(payload)
}

func (s *TestSuiteEnv) ResetPrimaryKey(tableName string) {
	s.api.DB.Exec("ALTER TABLE " + tableName + " DROP PRIMARY KEY, ADD PRIMARY KEY (  `id` )")
}

func (s *TestSuiteEnv) TestCreateBookOK() {
	s.ClearTable(&model.Book{})
	s.ResetPrimaryKey("Book")
	s.ClearTable(&model.Genre{})

	s.TestAddGenresOk()
	a := s.Assert()
	book := SampleBook

	reqBody, err := json.Marshal(book)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateBookRouter(s, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := model.Book{}
	if err := json.Unmarshal(body, &actual.ID); err != nil {
		a.Error(err)
	}

	expected := book
	a.Equal(expected.ID, actual.ID)
	s.ClearTable(&model.Book{})
}

func setCreateBookRouter(s *TestSuiteEnv, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	s.provider.Router = mux.NewRouter()

	s.provider.Post("/book", s.api.CreateBook)
	req, err := http.NewRequest(http.MethodPost, "/book", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.provider.Router.ServeHTTP(w, req)
	return req, w, nil
}

func (s *TestSuiteEnv) TestGetBookById() {
	a := s.Assert()
	book := SampleGetBookById
	err := insertTestBook(s, &book)
	if err != nil {
		a.Error(err)
	}

	bookId := utils.UintToString(SampleGetBookById.ID)
	req, w := setGetBookRouter(s, bookId)

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
	expected := book
	a.Equal(expected, actual)
}

func setGetBookRouter(s *TestSuiteEnv, bookId string) (*http.Request, *httptest.ResponseRecorder) {
	s.provider.Router = mux.NewRouter()
	s.provider.Get("/book/{id}", s.api.GetBookById)

	url := "/book/" + bookId
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": bookId,
	}

	req = mux.SetURLVars(req, vars)

	w := httptest.NewRecorder()
	s.provider.Router.ServeHTTP(w, req)
	return req, w
}

func insertTestBook(s *TestSuiteEnv, book *model.Book) error {
	err := s.api.DB.Create(&book).Error
	return err
}

func (s *TestSuiteEnv) TestUpdateBookOK() {
	s.ClearTable(&model.Book{})
	s.ResetPrimaryKey("Book")
	s.ClearTable(&model.Genre{})

	s.TestAddGenresOk()
	a := s.Assert()
	book := SampleUpdateBookOk

	err := insertTestBook(s, &SampleBook)
	if err != nil {
		a.Error(err)
	}

	reqBody, err := json.Marshal(book)
	if err != nil {
		a.Error(err)
	}
	bookId := utils.UintToString(SampleUpdateBookOk.ID)
	req, w, err := setUpdateBookRouter(s, bookId, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := model.Book{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := book
	a.Equal(expected, actual)
	s.ClearTable(&model.Book{})
}

func setUpdateBookRouter(s *TestSuiteEnv, bookId string, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	s.provider.Router = mux.NewRouter()

	s.provider.Put("/books/{id}", s.api.UpdateBook)
	url := "/books/" + bookId
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": bookId,
	}

	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()
	s.provider.Router.ServeHTTP(w, req)
	return req, w, nil
}

func (s *TestSuiteEnv) TestDeleteBookOk() {
	s.ClearTable(&model.Book{})
	s.ResetPrimaryKey("Book")
	s.ClearTable(&model.Genre{})

	s.TestAddGenresOk()

	a := s.Assert()
	book := SampleGetBookById
	err := insertTestBook(s, &book)
	if err != nil {
		a.Error(err)
	}

	bookId := utils.UintToString(SampleGetBookById.ID)
	req, w := setDeleteBookRouter(s, bookId)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
}

func setDeleteBookRouter(s *TestSuiteEnv, bookId string) (*http.Request, *httptest.ResponseRecorder) {
	s.provider.Router = mux.NewRouter()
	s.provider.Delete("/book/{id}", s.api.DeleteBook)

	url := "/book/" + bookId
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": bookId,
	}

	req = mux.SetURLVars(req, vars)

	w := httptest.NewRecorder()
	s.provider.Router.ServeHTTP(w, req)
	return req, w
}
