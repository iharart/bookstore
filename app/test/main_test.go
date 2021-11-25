package test

import (
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	"log"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {

	/*pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.BuildAndRun("db", "../../../scripts/Dockerfile", []string{"MYSQL_ROOT_PASSWORD=admin"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	/*resource, err := pool.Run("db", "latest", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}*/

	/*if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("admin:admin@(localhost:%s)/app", resource.GetPort("3306/tcp")))
		//db, err = sql.Open("mysql", fmt.Sprintf("admin:admin@tcp(db:3306)/app"))
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

	os.Exit(code)*/
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=admin",
		"MYSQL_DATABASE=app", "MYSQL_USER=admin", "MYSQL_ROOT_PASSWORD=admin"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("2222/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestResponds(t *testing.T) {

}
