package pages

import (
	"GoApiServer/model"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
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

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("could not connect to the database")
	}
	defer db.Close()

	var users []model.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

	//fmt.Fprintf(w, "Endpoint: Return All Users ")
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
