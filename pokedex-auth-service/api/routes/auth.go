package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/entities"
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/user"
)

type SignIn struct {
	Password string `json:"password" binding:"required" bson:"password"`
	Email    string `json:"email" binding:"required" bson:"email"`
}

func AuthRouter(app fiber.Router, service user.Service) {
	app.Post("/user", addUser(service))
	app.Post("/user/sign-in", signIn(service))
}

func addUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody entities.User

		err := c.BodyParser(&requestBody)

		if err != nil {
			return c.JSON(&fiber.Map{
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

func signIn(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody SignIn

		err := c.BodyParser(&requestBody)

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result, dberr := service.SignIn(requestBody.Email, requestBody.Password)

		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}
