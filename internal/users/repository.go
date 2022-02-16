package users

import (
	"podbilling/model"
)

type Repository interface {
	Create(user model.User) error
	GetByID(id uint) (model.User, error)
	GetByLogin(login string) (model.User, error)
	Update(user model.User) error
	Delete(id uint) error
	GetWorkers() ([]model.User, error)
}
