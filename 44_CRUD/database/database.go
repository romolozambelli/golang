package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" /// Connection Driver
)

func Connect() (*sql.DB, error) {

	stringConnectionData := "test-user:golang@/golang?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", stringConnectionData)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
