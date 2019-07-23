package main

import (
	"GoApiServer/model"
	"GoApiServer/router"
	"fmt"
)

func main() {

	fmt.Println("Server up")

	model.InitialMigration()

	router.HandleRequests()
	
}
