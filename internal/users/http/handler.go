package http

import (
	users "podbilling/internal/users"
)

type Handler struct {
	UseCase users.UseCase
}

func NewHandler(usecase users.UseCase) *Handler {
	return &Handler{
		UseCase: usecase,
	}
}
