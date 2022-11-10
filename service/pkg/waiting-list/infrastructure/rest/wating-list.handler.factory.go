package rest

import (
	"github/abraxas-365/Leak/pkg/waiting-list/application"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	FollowACrush(c *fiber.Ctx) error
	DeleteAFollow(c *fiber.Ctx) error
	ShowWaitingList(c *fiber.Ctx) error
}

type handler struct {
	app application.Application
}

func HandlerFabric(app application.Application) Handler {
	return &handler{
		app,
	}
}
