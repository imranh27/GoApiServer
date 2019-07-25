package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error
var dbdialect = "sqlite3"
var dbname = "test.db"


type User struct {
	gorm.Model
	Name string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open(dbdialect, dbname)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func DbConnect() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("could not connect to the database")
	}
	defer db.Close()

}