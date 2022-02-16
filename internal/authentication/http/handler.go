package http

import (
	auth "podbilling/internal/authentication"
)

type Handler struct {
	UseCase auth.UseCase
}

func NewHandler(usecase auth.UseCase) *Handler {
	return &Handler{
		UseCase: usecase,
	}
}
