package authentication

import (
	"podbilling/authentication/model"
)

type Repository interface {
	GetUser(login string, password uint32) (model.User, error)
}
