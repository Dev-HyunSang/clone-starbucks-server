package middleware

import (
	"github.com/dev-hyunsang/clone-stackbuck-backend/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	users := api.Group("/users")
	users.Post("/signup", cmd.SignUpUserHandler)
}
