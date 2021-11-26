package main

import (
	"fmt"
	database "github.com/iharart/bookstore/app/database"
	router "github.com/iharart/bookstore/app/router"
	"log"
	"net/http"
)

const (
	PortNumber string = "1234"
)

func main() {
	database.Initialize()
	r := router.SetUp()
	fmt.Println("Connected to port " + PortNumber)
	log.Fatal(http.ListenAndServe(":"+PortNumber, r))
}
