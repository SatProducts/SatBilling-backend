package users

import (
	"podbilling/users/model"
)

type UseCase interface {
	Create(login, password string, permissons uint8) error
	Get(id uint) (model.User, error)
	Update(user model.User) error
	Delete(id uint) error
	GetWorkers() ([]model.User, error)
}
