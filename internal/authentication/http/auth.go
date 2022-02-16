package http

import (
	fiber "github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {

	var info AuthRequest

	if err := c.BodyParser(&info); err != nil {
		return fiber.ErrBadRequest
	}

	token, err := h.UseCase.GenerateJWT(c.UserContext(), info.Login, info.Password)

	if err != nil {
		return fiber.ErrUnauthorized
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
