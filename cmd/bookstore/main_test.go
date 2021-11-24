package main

import (
	handler "bookstore/handler"
	model "bookstore/model"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestSuiteEnv struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	service handler.Service
	book    *model.Book
	genre   *model.Genre
}

func NewDatabase() (*gorm.DB, sqlmock.Sqlmock) {

	// get db and mock
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}
	defer sqlDB.Close()

	dialector := mysql.New(mysql.Config{
		Conn:       sqlDB,
		DriverName: "mysql",
	})

	columns := []string{"version"}
	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
		mock.NewRows(columns).FromCSVString("1"),
	)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	return db, mock
}

func (s *TestSuiteEnv) SetupSuite() {
	db, mock := NewDatabase()
	if mock != nil {
		s.mock = mock
	} else {
		s.FailNow("mock is nil")
	}
	selectQ := "SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE '%' ORDER BY SCHEMA_NAME='' DESC limit 1"
	creatingQ := "CREATE TABLE `Genre` (`id` bigint AUTO_INCREMENT,`name` longtext,PRIMARY KEY (`id`))"

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(selectQ)
	s.mock.ExpectClose()
	s.mock.ExpectBegin()
	s.mock.ExpectExec(creatingQ)
	s.mock.ExpectBegin()
	s.service.DB = model.Migrate(db)
	s.mock.ExpectClose()
	s.mock.ExpectCommit()

	if err := mock.ExpectationsWereMet(); err != nil {

		s.Errorf(err, "there were unfulfilled expectations")
	}
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(TestSuiteEnv))
}

func (s *TestSuiteEnv) TestGetBooksEmptyResult() {
	//books := []model.Book{}
	s.mock.ExpectQuery("SELECT * FROM `Book` ORDER BY name").WillReturnRows(sqlmock.NewRows([]string{""}))
	//s.mock.ExpectQuery((s.service.DB.Debug().Preload(model.GENRE).Order("name").Find(&books)))
	req, w := setGetBooksRouter(&s.service)
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

func setGetBooksRouter(service *handler.Service) (*http.Request, *httptest.ResponseRecorder) {
	//service := &handler.Service{DB: db}
	service.Router = mux.NewRouter()
	service.Get("/books", service.GetAllBooks)
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	service.Router.ServeHTTP(w, req)
	return req, w
}

/*var db *sql.DB
func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	/*resource, err := pool.BuildAndRun("mysql", "../../scripts/Dockerfile", []string{"MYSQL_ROOT_PASSWORD=admin"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}*/
/*resource, err := pool.Run("db", "latest", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("admin:admin@(localhost:%s)/bookstore", resource.GetPort("3306/tcp")))
		//db, err = sql.Open("mysql", fmt.Sprintf("admin:admin@tcp(db:3306)/bookstore"))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}*/

/*func TestResponds(t *testing.T) {

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	opts := dockertest.RunOptions{
		Repository:   "db",
		Hostname:      "db",

		Tag:          "latest",
		Env:          []string{"MYSQL_ROOT_PASSWORD=admin"},
		User:          "admin",
		ExposedPorts: []string{"3406"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3406": {
				{HostIP: "0.0.0.0", HostPort: "3406"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	var resp *http.Response

	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("8082/tcp"), "/books"))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement!
	require.Contains(t, string(body), "<3", "does not respond with love?")
}*/

/*func TestGetBooksEmptyResult(t *testing.T) {
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
}*/
