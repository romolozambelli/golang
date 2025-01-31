package main

import (
	"log"
	"net/http"
)

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Users with func"))

}

func main() {
	// HTTP is a communication protocol - Cliente Server

	// Routes/Address defines what to do
	// URI --> Identify the resource
	// Method --> POST / GET / PUT / DELETE

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
