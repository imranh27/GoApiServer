package router

import (
	"GoApiServer/pages"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", pages.HomePage).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", pages.UserCreate).Methods("POST")
	myRouter.HandleFunc("/users", pages.UsersReadAll).Methods("GET")
	myRouter.HandleFunc("/users/{name}", pages.UserReadOne).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", pages.UserUpdateEmail).Methods("PUT")
	myRouter.HandleFunc("/users/{name}", pages.UserDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

