package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Person Struct (Model)
type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Age       string `json:"age,omitempty"`
	Address   string `json:"address,omitempty"`
}

//Init persons var as a slice Person struct
var persons []Person

// Get All Persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

// Get Single Person by Id
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through person and find with id
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Get Single Person by FirstName
func getPersonByFirstName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through person and find with fistname
	for _, itemf := range persons {
		if itemf.Firstname == params["firstname"] {
			json.NewEncoder(w).Encode(itemf)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Get Single Person by age
func getPersonByAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through person and find with id
	for _, item := range persons {
		if item.Age == params["age"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Create a New Person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)

}

// Update Person
func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = params["id"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	json.NewEncoder(w).Encode(persons)

}

// Delete Person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello  Julius</h1>")
}
func main() {
	// Init Router
	r := mux.NewRouter()
	InitiaMigration()

	// Mock Person Data
	persons = append(persons, Person{ID: "100", Firstname: "Mark", Lastname: "Slade", Address: "200 E St", Age: "20"})
	persons = append(persons, Person{ID: "101", Firstname: "John", Lastname: "Luski", Address: "100 L St", Age: "30"})
	persons = append(persons, Person{ID: "102", Firstname: "Lux", Lastname: "Blake", Address: "300 D St", Age: "19"})

	// Route Handlers / Endpoint for Person API
	r.HandleFunc("/api/persons", getPersons).Methods("GET")
	r.HandleFunc("/api/persons/{id}", getPerson).Methods("GET")
	r.HandleFunc("/api/persons/{firstname}", getPersonByFirstName).Methods("GET")
	r.HandleFunc("/api/persons/{age}", getPersonByAge).Methods("GET")
	r.HandleFunc("/api/persons", createPerson).Methods("POST")
	r.HandleFunc("/api/persons/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/api/persons/{id}", deletePerson).Methods("DELETE")
	http.HandleFunc("/", index)

	fmt.Println("server starting...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
