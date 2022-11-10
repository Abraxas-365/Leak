package rest

import (
	"github/abraxas-365/Leak/pkg/auth/application"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CallBackInstagram(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type handler struct {
	app application.Application
}

func HandlerFabric(app application.Application) Handler {
	return &handler{
		app: app,
	}
}
