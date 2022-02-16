package users

import (
	"errors"
)

var (
	UserAlreadyExistsError = errors.New("user already exists")
	EmptyFieldError        = errors.New("empty field")
	NoPermissionsError     = errors.New("no permissions to perform operation")
)
