package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/raphaelcoutu/go-azure-rhpharma/database"
	"github.com/raphaelcoutu/go-azure-rhpharma/routes"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Azure RHPharma")
	})

	routes.Register(app)

	log.Fatal(app.Listen("localhost:3000"))
}
