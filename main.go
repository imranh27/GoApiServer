package main

import (
	"apitest/router"
	"fmt"
)

func main() {

	fmt.Println("Server up")
	router.HandleRequests()
	
}
