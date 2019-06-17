package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.db
var err error

type Person struct {
	gorm.Model
	Id        string
	Firstname string
	Lastname  string
	Age       string
	Address   string
}

func InitiaMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nill {
		fmt.Println(err.Erroe())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Person{})
}

func FindAllPersons(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("Unable to connect to the database")
	}
	defer db.close()

	var persons []Person
	db.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}


func NewPersonCreate(w http.ResponseWriter, r *http.Request){
db, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("Unable to connect to the database")
	}
	defer db.close()


	vars := mux.Vars(r)
	firstname := vars("firstname")
	lastname := vars("lastname")
	age      := vars("age")
	address  := vars("address")

	db.Create(&Person{Firstname: firstname, Lastname: lastname, Age: age, Address: address})

	fmt.Fprintf(w, "New Person Created")
}

	func FindPersonByFirstName(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("Unable to connect to the database")
	}
	defer db.close()

	vars := mux.Vars(r)
	firstname := vars["firstname"]

	var person Person
	db.Where("firstname = ?", firstname).Find($person)


}

    func FindPersonByAge(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("Unable to connect to the database")
	}
	defer db.close()

	vars := mux.Vars(r)
	firstname := vars["age"]

	var person Person
	db.Where("age = ?", age).Find($person)


}

    

    func DeletePersonByFirstName(w http.ResponseWriter, r *http.Request){
		db, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("Unable to connect to the database")
	}
	defer db.close()

	vars := mux.Vars(r)
	firstname := vars["firstname"]

	var person Person
	db.Where("firstname = ?", firstname).Find($person)
	db.Delete(&person)

	fmt.Fprintf(w, "Person Deleted")







}