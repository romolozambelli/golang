package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// CREATE = POST / READ = GET / UPDATE = PUT / DELETE = DELETE

	router := mux.NewRouter()

	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)        // Just a embedded string of POST
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)           // Just a embedded string of GET for all the users
	router.HandleFunc("/users/{id}", server.GetUser).Methods(http.MethodGet)       // Just a embedded string of GET single user
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)    // Just a embedded string of PUT single user
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete) // Just a embedded string of DELETE single user

	fmt.Println("Listenning on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
