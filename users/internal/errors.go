package users

import (
	"errors"
)

var (
	UserAlreadyExistsError = errors.New("User already exists")
	EmptyFieldError = errors.New("Empty field")
	NoPermissionsError = errors.New("No permissions to perform operation")
)