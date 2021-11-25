package handler

import (
	"database/sql"
	"fmt"
	"github.com/iharart/bookstore/app/configs"
	"github.com/iharart/bookstore/app/model"
	"github.com/iharart/bookstore/app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func (s *Service) Initialize() {
	connectionString := configs.GetConnectionString()
	fmt.Println(connectionString)

	var sqlDb *sql.DB
	var err error

	count := utils.StringToInt(os.Getenv("MYSQL_MAX_CONN_COUNT"))
	timeout := utils.StringToInt(os.Getenv("MYSQL_TIMEOUT"))

	for i := 0; i <= count; i++ {
		sqlDb, err = sql.Open("mysql", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		if i == 2 {
			fmt.Println("Waiting opening database")
			break
		}
		if err := sqlDb.Ping(); err != nil {
			fmt.Printf("Conection trying %d\n", i)
			time.Sleep(time.Duration(timeout) * time.Second)
		} else {
			break
			fmt.Println("Ping succeeded")
		}
	}
	fmt.Println("Successfully connect to sqlDb")

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	s.DB = model.Migrate(db)
}
