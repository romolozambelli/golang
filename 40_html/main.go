package main

import (
	"log"
	"net/http"
)

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Users with func"))

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from Home"))

}

func html(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from Home"))

}

func main() {
	// HTTP is a communication protocol - Cliente Server

	// Routes/Address defines what to do
	// URI --> Identify the resource
	// Method --> POST / GET / PUT / DELETE

	http.HandleFunc("/home", users)

	http.HandleFunc("/users", users)

	http.HandleFunc("/html", html)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
