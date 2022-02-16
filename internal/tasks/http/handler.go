package http

import (
	"podbilling/internal/tasks"
)

type Handler struct {
	UseCase tasks.UseCase
}

func NewHandler(uc tasks.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}
