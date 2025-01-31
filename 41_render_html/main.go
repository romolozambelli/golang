package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Users with func"))

}

func home(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("*.html"))

	u1 := user{"John", "john.wick@example.com"}

	templates.ExecuteTemplate(w, "home.html", u1)

}

type user struct {
	Name  string
	Email string
}

func main() {

	fmt.Println("Begin of Application")

	// HTTP is a communication protocol - Cliente Server

	// Routes/Address defines what to do
	// URI --> Identify the resource
	// Method --> POST / GET / PUT / DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
