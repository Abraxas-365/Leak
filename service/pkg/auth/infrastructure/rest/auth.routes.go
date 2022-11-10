package rest

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler Handler) {
	auth := app.Group("/auth")

	auth.Post("/register", handler.CallBackInstagram)
	auth.Post("/login", handler.Login)
}
