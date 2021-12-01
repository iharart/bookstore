package test

import (
	"database/sql"
	"github.com/iharart/bookstore/app/handler"
	"github.com/iharart/bookstore/app/router"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type TestSuiteEnv struct {
	suite.Suite
	pool     *dockertest.Pool
	resource *dockertest.Resource
	sqlDb    *sql.DB
	api      handler.APIEnv
	provider router.Provider
}
