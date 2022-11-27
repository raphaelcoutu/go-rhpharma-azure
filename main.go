package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB
var server = ""
var port = 1433
var user = ""
var password = ""
var database = ""

type User struct {
	Id        uint   `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Azure RHPharma")
	})

	app.Get("/token", GetToken)
	app.Get("/users", GetUsers)
	app.Get("/users/:userId<int>", GetUser)
	app.Get("/constraints", GetConstraints)
	app.Get("/constraintTypes", GetConstraintTypes)
	app.Get("/constraintTypes/:typeId<int>", GetConstraintType)

	log.Fatal(app.Listen(":3000"))
}
