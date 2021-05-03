package main

import (
	"micro-echo/router"
)

//main function
func main() {

	// create a new echo instance
	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
