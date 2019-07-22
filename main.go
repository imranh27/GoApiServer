package main

import (
	"GoApiServer/router"
	"fmt"
)

func main() {

	fmt.Println("Server up")

	router.HandleRequests()
	
}
