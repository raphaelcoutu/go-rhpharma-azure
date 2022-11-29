package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/token", GetToken)
	app.Get("/users", Protected, GetUsers)
	app.Get("/users/:userId<int>", Protected, GetUser)
	app.Get("/constraints", Protected, GetConstraints)
	app.Get("/constraintTypes", GetConstraintTypes)
	app.Get("/constraintTypes/:typeId<int>", GetConstraintType)
	app.Get("/stats", GetDBStats)
}

func Protected(c *fiber.Ctx) error {
	qToken := c.Query("token")
	if qToken == "" {
		return c.SendStatus(403)
	}

	if ValidateToken(qToken) {
		return c.Next()
	}

	return c.SendStatus(403)
}
