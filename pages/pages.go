package pages

import (
	"GoApiServer/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

var dbdialect = "sqlite3"
var dbname = "test.db"

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint: homepage")

}

//Create
func UserCreate(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&model.User{Name: name, Email: email})

	fmt.Fprintf(w, "Created User")
}

//Read
func UsersReadAll(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(dbdialect, dbname)
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

	db, err := gorm.Open(dbdialect, dbname)
	if err != nil {
		panic("could not connect to database")
	}
	defer db.Close()

	fmt.Fprintf(w, "Endpoint: Return a User")
}

//Update
func UserUpdateEmail(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbdialect,dbname)
	if err != nil {
		panic("could not connect to database")
	}
	defer db.Close()

	fmt.Fprintf(w, "Endpoint: Update User")
}

//Delete
func UserDelete(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbdialect, dbname)
	if err != nil {
		panic("could not open databaase")
	}
	defer db.Close()

	fmt.Fprintf(w, "Endpoint: Delete User")
}
