package test

import (
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
	wrapper  router.Wrapper
	book     *model.Book
	genre    *model.Genre
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
		config.RestartPolicy = docker.RestartUnlessStopped()
	})
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		s.sqlDb, err = sql.Open("mysql", fmt.Sprintf("admin:admin@(localhost:%s)/bookstore", resource.GetPort("3306/tcp")))
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
	if err != nil {
		log.Fatal(err)
	}
	s.api.DB = database.Migrate(db)
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
	s.wrapper.Router = mux.NewRouter()
	s.wrapper.Get("/books", s.api.GetAllBooks)
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.wrapper.Router.ServeHTTP(w, req)
	return req, w
}

func (s *TestSuiteEnv) ClearTable(payload interface{}) {
	s.api.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(payload)
}
