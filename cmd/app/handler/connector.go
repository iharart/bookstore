package handler

import (
	"bookstore/configs"
	"bookstore/model"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Service struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (s *Service) Initialize() {
	connectionString := configs.GetConnectionString()
	fmt.Println(connectionString)

	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start working with GORM..")
	s.DB = model.Migrate(db)
	s.Router = mux.NewRouter()
	s.SetUpRouters()
}
