package authentication

import (
	"podbilling/model"
)

type Repository interface {
	Get(login, password string) (model.User, error)
}
