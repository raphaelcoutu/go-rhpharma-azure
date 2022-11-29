package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raphaelcoutu/go-rhpharma-azure/database"
	"github.com/raphaelcoutu/go-rhpharma-azure/logger"
)

type ConstraintType struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetConstraintTypes(c *fiber.Ctx) error {
	ctx := context.Background()
	db := database.DB

	err := db.PingContext(ctx)
	if err != nil {
		logger.Log("Error: PingContext. Message: " + err.Error())
		return err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Description FROM ConstraintTypes;")

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		logger.Log("Error: QueryContext. Message: " + err.Error())
		return err
	}

	defer rows.Close()

	types := []ConstraintType{}

	for rows.Next() {
		var t ConstraintType
		var description sql.NullString
		err := rows.Scan(&t.Id, &t.Name, &description)
		if err != nil {
			logger.Log("Error: Scan. Message: " + err.Error())
			return err
		}

		if description.Valid {
			t.Description = description.String
		}

		types = append(types, t)
	}

	jsonData, _ := json.Marshal(types)
	return c.SendString(string(jsonData))
}

func GetConstraintType(c *fiber.Ctx) error {
	ctx := context.Background()
	db := database.DB

	err := db.PingContext(ctx)
	if err != nil {
		logger.Log("Error: PingContext. Message: " + err.Error())
		return err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Description FROM ConstraintTypes Where Id = @Id;")

	stmt, err := db.Prepare(tsql)
	if err != nil {
		logger.Log("Error: Prepare. Message: " + err.Error())
		return err
	}
	defer stmt.Close()

	typeId, err := c.ParamsInt("typeId")
	if err != nil {
		logger.Log("Error: ParamsInt. Message: " + err.Error())
		return err
	}

	row := stmt.QueryRowContext(ctx, sql.Named("Id", typeId))

	var t ConstraintType
	var description sql.NullString
	err = row.Scan(&t.Id, &t.Name, &description)
	if err != nil {
		logger.Log("Error: Scan. Message: " + err.Error())
		return err
	}

	if description.Valid {
		t.Description = description.String
	}

	jsonData, _ := json.Marshal(t)
	return c.SendString(string(jsonData))
}
