package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/entities"
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/user"
)

func AuthRouter(app fiber.Router, service user.Service) {
	app.Post("/user", addUser(service))
}

func addUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody entities.User

		err := c.BodyParser(&requestBody)

		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result, dberr := service.InsertUser(&requestBody)

		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}
