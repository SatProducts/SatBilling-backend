package http

import (
	"podbilling/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func (h *Handler) Get(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.ErrNotFound
	}

	taskModel, err := h.UseCase.Get(c.UserContext(), (uint(id)))

	if err != nil {
		return fiber.ErrInternalServerError
	}

	userID := uint(user["id"].(float64))
	permissions := uint8(user["permissions"].(float64))

	if userID != taskModel.FromUser && userID != taskModel.ForUser && permissions != 3 {
		return fiber.ErrForbidden
	}

	return c.JSON(taskModel)
}

func (h *Handler) Create(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	if permissions != model.OPERATOR {
		return fiber.ErrForbidden
	}

	var taskModel model.Task

	if err := c.BodyParser(&taskModel); err != nil {
		return fiber.ErrBadRequest
	}

	taskModel.FromUser = uint(user["id"].(float64))

	if err := h.UseCase.Create(c.UserContext(), taskModel); err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusCreated)
	return nil
}

func (h *Handler) Update(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.ErrNotFound
	}

	taskModel, err := h.UseCase.Get(c.UserContext(), uint(id))

	if err != nil {
		return fiber.ErrNotFound
	}

	if uint(user["id"].(float64)) != taskModel.FromUser {
		return fiber.ErrForbidden
	}

	if err := c.BodyParser(&taskModel); err != nil {
		return fiber.ErrBadRequest
	}

	if err = h.UseCase.Update(c.UserContext(), taskModel); err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return nil
}

func (h *Handler) Delete(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	permissions := uint8(user["permissions"].(float64))

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.ErrNotFound
	}

	taskModel, err := h.UseCase.Get(c.UserContext(), uint(id))

	if err != nil {
		return fiber.ErrNotFound
	}

	if user["id"] != taskModel.FromUser && permissions != model.ADMINISTRATOR {
		return fiber.ErrForbidden
	}

	if err = h.UseCase.Delete(c.UserContext(), taskModel.ID); err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return nil
}

func (h *Handler) GetFor(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	tasks, err := h.UseCase.GetFor(c.UserContext(), user["id"].(uint))

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(tasks)
}

func (h *Handler) GetFrom(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	tasks, err := h.UseCase.GetFrom(c.UserContext(), user["id"].(uint))

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(tasks)
}
