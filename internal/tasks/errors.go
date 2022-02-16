package tasks

import (
	"errors"
)

var (
	TaskNotFoundError  = errors.New("task not found")
	NoWorkersError     = errors.New("no workers in the database")
	EmptyFieldError    = errors.New("empty field in model")
	NoPermissionsError = errors.New("no permissions to perform operation")
)
