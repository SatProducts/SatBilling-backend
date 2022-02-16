package model

import (
	"errors"
)

var (
	DeleteAdminError = errors.New("cannot delete an admin user")
)