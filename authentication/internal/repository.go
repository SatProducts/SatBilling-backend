package authentication

import (
	"podbilling/authentication/model"
)

type Repository interface {
	GetUser(login, password string) (model.User, error)
}
