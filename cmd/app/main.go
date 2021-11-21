package main

import (
	handler "bookstore/handler"
	"fmt"
)

const (
	PortNumber string = "1234"
)

func Initialize() {

	service := &handler.App{}
	service.Initialize()
	fmt.Println("Connected to port " + PortNumber)
	service.Run(":" + PortNumber)
}

func main() {
	Initialize()
}
