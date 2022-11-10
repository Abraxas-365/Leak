package rest

import (
	"github/abraxas-365/Leak/internal/jwtauth"
	"github/abraxas-365/Leak/internal/leakerrs"
	"github/abraxas-365/Leak/pkg/waiting-list/domain/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
FollowACrush
*/
func (h *handler) FollowACrush(c *fiber.Ctx) error {
	follow := models.Follow{}
	if err := c.BodyParser(&follow); err != nil {
		err := leakerrs.GetError(err)
		c.Status(err.Code).JSON(err)
	}
	return c.SendStatus(201)
}

/*
DeleteAFollow
param=CrushId
*/
func (h *handler) DeleteAFollow(c *fiber.Ctx) error {
	user, err := jwtauth.ExtractTokenMetadata(c)
	if err != nil {
		err := leakerrs.GetError(err)
		c.Status(err.Code).JSON(err)
	}

	crushId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		err := leakerrs.GetError(err)
		c.Status(err.Code).JSON(err)
	}

	if err := h.app.DeleteAFollow(crushId, user.ID); err != nil {
		err := leakerrs.GetError(err)
		c.Status(err.Code).JSON(err)
	}

	return c.SendStatus(201)
}

/*Aun no implementar, definir si es algo primiun o no*/
func (h *handler) ShowWaitingList(c *fiber.Ctx) error
