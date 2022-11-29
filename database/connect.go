package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB
var server = os.Getenv("DB_HOST")
var database = os.Getenv("DB_DATABASE")
var port = 1433
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASS")

func ConnectDB() {
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err.Error())
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)

	fmt.Println("Database: connection established.")

}
