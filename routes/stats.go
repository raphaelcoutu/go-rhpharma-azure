package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/raphaelcoutu/go-rhpharma-azure/database"
)

func GetDBStats(c *fiber.Ctx) error {
	stats, _ := json.Marshal(database.DB.Stats())
	return c.SendString(string(stats))
}
