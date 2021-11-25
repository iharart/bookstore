package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBConfig struct {
	User       string
	Password   string
	ServerName string
	DBPort     string
	DBName     string
}

func getConfig() *DBConfig {
	return &DBConfig{
		User:       os.Getenv("MYSQL_USER"),
		Password:   os.Getenv("MYSQL_ROOT_PASSWORD"),
		ServerName: os.Getenv("MYSQL_SERVERNAME"),
		DBPort:     os.Getenv("MYSQL_PORT"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
	}
}

func GetConnectionString() string {
	config := getConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User,
		config.Password, config.ServerName, config.DBPort, config.DBName)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
