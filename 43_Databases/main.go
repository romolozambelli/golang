package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Users to connect on database
	stringConnectionData := "test-user:golang@/golang?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", stringConnectionData)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close() // avoid open connections

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Open Connection")

	line, erro := db.Query("select * from users")

	if erro != nil {
		log.Fatal(erro)
	}
	defer line.Close()

	fmt.Println(line)

}
