package pages

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: homepage")

}

//Create
func UserCreate(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: Create User")
}

//Read
func UsersReadAll(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: Return All Users ")

}

func UserReadOne(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: Return a User")
}

//Update
func UserUpdateEmail(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: Update User")
}

//Delete
func UserDelete(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: Delete User")
}
