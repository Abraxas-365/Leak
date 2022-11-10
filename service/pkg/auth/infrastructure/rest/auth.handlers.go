package rest

import (
	"fmt"
	"github/abraxas-365/Leak/internal/leakerrs"
	"github/abraxas-365/Leak/pkg/auth/domain/models"

	"github.com/gofiber/fiber/v2"
)

/*
this will take from param the token from ig
for test porpus Im sending the token in the post body
*/
func (h *handler) CallBackInstagram(c *fiber.Ctx) error {
	register := models.RegisterWithPlatform{}

	if err := c.BodyParser(&register); err != nil {
		err := leakerrs.GetError(err)
		c.Status(err.Code).JSON(err)
	}
	token, err := h.app.RegisterWithInstagram(register)
	fmt.Println(err)
	if err != nil {
		err := leakerrs.GetError(err)
		return c.Status(err.Code).JSON(err)
	}

	return c.Status(201).JSON(token)
}
