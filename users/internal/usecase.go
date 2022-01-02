package users

import (
	"podbilling/users/model"
)

type UseCase interface {
	CreateUser(login, password string, permissons uint8) error
	GetUser(id uint) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id uint) error
	GetAllWorkers() ([]model.User, error)
}
