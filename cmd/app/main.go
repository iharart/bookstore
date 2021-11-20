package main

import (
	handler "bookstore/handler"
	"fmt"
)

func Initialize() {

	service := &handler.App{}
	service.Initialize()
	fmt.Println("Connected to port 1234")
	service.Run(":1234")
}

func main() {
	Initialize()
}
