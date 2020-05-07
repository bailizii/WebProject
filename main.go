package main

import (
	"WebProject/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	handlers.AddData()

	router.HandleFunc("/people",handlers.CreatePerson ).Methods("POST")
	router.HandleFunc("/people", handlers.GetAllPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")

	fmt.Println("Starting server on port 5000...")
	http.ListenAndServe(":5000", router)
}