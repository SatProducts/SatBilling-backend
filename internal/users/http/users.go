package http

import (
	"podbilling/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) GetSelf(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	userModel, err := h.UseCase.Get(c.UserContext(), uint(user["id"].(float64)))

	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(userModel)
}

func (h *Handler) Get(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	if permissions != model.ADMINISTRATOR {
		return fiber.ErrForbidden
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.ErrNotFound
	}

	userModel, err := h.UseCase.Get(c.UserContext(), uint(id))

	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(userModel)
}

func (h *Handler) Create(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	if permissions != model.ADMINISTRATOR {
		return fiber.ErrForbidden
	}

	var newUser model.User

	if err := c.BodyParser(&newUser); err != nil {
		return fiber.ErrBadRequest
	}

	err := h.UseCase.Create(
		c.UserContext(),
		newUser,
	)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusCreated)
	return nil
}

func (h *Handler) Update(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	if permissions != model.ADMINISTRATOR {
		return fiber.ErrForbidden
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.ErrNotFound
	}

	userModel, err := h.UseCase.Get(c.UserContext(), uint(id))

	if err != nil {
		return fiber.ErrNotFound
	}

	if err = c.BodyParser(&userModel); err != nil {
		return fiber.ErrBadRequest
	}

	if err = h.UseCase.Update(c.UserContext(), userModel); err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return nil
}

func (h *Handler) Delete(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	if permissions != model.ADMINISTRATOR {
		return fiber.ErrForbidden
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.ErrNotFound
	}

	_, err = h.UseCase.Get(c.UserContext(), uint(id))

	if err != nil {
		return fiber.ErrNotFound
	}

	if err = h.UseCase.Delete(c.UserContext(), uint(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return nil
}

func (h *Handler) GetWorkers(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	if permissions != model.ADMINISTRATOR {
		return fiber.ErrForbidden
	}

	workers, err := h.UseCase.GetWorkers(c.UserContext())

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(workers)
}
