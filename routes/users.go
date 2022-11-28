package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raphaelcoutu/go-rhpharma-azure/database"
)

type User struct {
	Id        uint   `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

func GetUsers(c *fiber.Ctx) error {
	ctx := context.Background()
	db := database.DB

	err := db.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := fmt.Sprintf("SELECT Id, LastName, FirstName FROM Users;")

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User

		err := rows.Scan(&u.Id, &u.LastName, &u.FirstName)
		if err != nil {
			return err
		}

		users = append(users, u)
	}

	jsonData, _ := json.Marshal(users)
	return c.SendString(string(jsonData))
}

func GetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	db := database.DB

	err := db.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := fmt.Sprintf("SELECT Id, LastName, FirstName FROM Users Where Id = @Id;")

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	userId, err := c.ParamsInt("userId")
	if err != nil {
		return err
	}

	row := stmt.QueryRowContext(ctx, sql.Named("Id", userId))

	var u User
	err = row.Scan(&u.Id, &u.LastName, &u.FirstName)
	if err != nil {
		return err
	}

	jsonData, _ := json.Marshal(u)
	return c.SendString(string(jsonData))
}
