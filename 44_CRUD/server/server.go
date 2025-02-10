package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

// Inser users into the Database
func CreateUser(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Body does not match, wrong data"))
		return
	}

	var user user

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte("Wrong user format"))
		return
	}

	fmt.Println(user)

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Wrong user format"))
		return
	}

	statement, erro := db.Prepare("insert into users (name, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Fail to create statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)

	if erro != nil {
		w.Write([]byte("Error to execute the statement"))
		return
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Error to retrieve the id of the inserted users"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted with success on ID = %d", idInsert)))

}

// Get Users on the database
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Error Connecting to database"))
		return
	}
	defer db.Close()

	lines, erro := db.Query("SELECT * FROM users")
	if erro != nil {
		w.Write([]byte("Error reading users from database"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user

		if erro := lines.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Error on scanning users"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Error on converting users to JSON"))
		return
	}

}

// Get a single user on the database
func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Error converting values to int"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Error Connecting to database"))
		return
	}
	defer db.Close()

	line, erro := db.Query("SELECT * FROM users WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Error reading user from database"))
		return
	}
	defer line.Close()

	var user user

	if line.Next() {
		if erro := line.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Error on scanning users"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Error on converting user to JSON"))
		return
	}

}

// Update users data on the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Error converting values to int"))
		return
	}

	// Read the body data
	bodyRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Body does not match, wrong data"))
		return
	}

	var user user
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte("Wrong user format"))
		return
	}

	// Connect to database
	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Error Connecting to database"))
		return
	}
	defer db.Close()

	// Update Users
	statement, erro := db.Prepare("UPDATE users SET name = ?, email = ? where id = ?")

	if erro != nil {
		w.Write([]byte("Fail to create statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, ID); erro != nil {
		w.Write([]byte("Error to execute the statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Error converting values to int"))
		return
	}

	// Connect to database
	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Error Connecting to database"))
		return
	}
	defer db.Close()

	// Update Users
	statement, erro := db.Prepare("DELETE FROM users WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Fail to create statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Error to execute the statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
