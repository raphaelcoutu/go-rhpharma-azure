package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/raphaelcoutu/go-rhpharma-azure/logger"
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
		logger.Log("Error: sql.Open(). Message: " + err.Error())
		log.Fatal(err.Error())
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)

	ctx := context.Background()

	err = DB.PingContext(ctx)
	if err != nil {
		logger.Log("Error: PingContext. Message: " + err.Error())
	}

	fmt.Println("Database: connection established.")

}
