package users

import (
	"podbilling/users/model"
)

type Repository interface {
	Create(login, password string, permissons uint8) error
	GetByID(id uint) (model.User, error)
	GetByLogin(login string) (model.User, error)
	Update(user model.User) error
	Delete(id uint) error
	GetWorkers() ([]model.User, error)
}
