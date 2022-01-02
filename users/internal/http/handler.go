package http

import (
	users "podbilling/users/internal"
)

type UsersHandler struct {
	UseCase users.UseCase
}

func NewUsersHandler(usecase users.UseCase) *UsersHandler {
	return &UsersHandler{
		UseCase: usecase,
	}
}