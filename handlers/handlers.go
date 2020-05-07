package handlers

import (
	"WebProject/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var people []model.Person = []model.Person{}

func AddData() {
	people = append(people, model.Person{Firstname: "San", Lastname: "Zhang", Phone:"18301236711"})
	people = append(people, model.Person{Firstname: "Wu",	Lastname: "Wang", Phone:"16201236711"})
}


func GetPerson(w http.ResponseWriter, r *http.Request) {
	// get the ID of the person from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(people) {
		w.WriteHeader(404)
		w.Write([]byte("Nothing found with specified ID"))
		return
	}

	person := people[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPerson model.Person
	json.NewDecoder(r.Body).Decode(&newPerson)

	people = append(people, newPerson)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(people)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	// get the ID of the person from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(people) {
		w.WriteHeader(404)
		w.Write([]byte("Nothing found with specified ID"))
		return
	}

	// get the value from JSON body
	var updatedPerson model.Person
	json.NewDecoder(r.Body).Decode(&updatedPerson)

	people[id] = updatedPerson

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPerson)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// get the ID of the person from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(people) {
		w.WriteHeader(404)
		w.Write([]byte("Nothing found with specified ID"))
		return
	}

	// Delete the post from the slice
	// https://github.com/golang/go/wiki/SliceTricks#delete
	people = append(people[:id], people[id+1:]...)

	w.WriteHeader(200)
}