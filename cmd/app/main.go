package main

import (
	handler "bookstore/handler"
	"fmt"
)

func Initialize() {

	app := &handler.App{}
	app.Initialize()
	fmt.Println("Connected to port 1234")
	app.Run(":1234")

}
func main() {
	Initialize()
}
