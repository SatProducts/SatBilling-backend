package http

import (
	auth "podbilling/authentication/internal"
)

type AuthHandler struct {
	UseCase auth.UseCase
}

func NewAuthHandler(usecase auth.UseCase) *AuthHandler {
	return &AuthHandler{
		UseCase: usecase,
	}
}
