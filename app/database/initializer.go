package database

import (
	"database/sql"
	"fmt"
	"github.com/iharart/bookstore/app/configs"
	"github.com/iharart/bookstore/app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Initialize() {
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
			fmt.Println("Ping succeeded")
			break
		}
	}
	fmt.Println("Successfully connect to sqlDb")

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.ErrorCheck(err)

	DB, err = Migrate(db)
	utils.ErrorCheck(err)
}

func GetDB() *gorm.DB {
	return DB
}
