package rest

import (
	"github/abraxas-365/Leak/internal/jwtauth"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler Handler) {
	auth := app.Group("/waitinglist")

	auth.Post("/", handler.FollowACrush)
	auth.Delete("/:id", jwtauth.JWTProtected(), handler.DeleteAFollow)
	// auth.Get("/:id", handler.ShowWaitingList)
}
