package users

import (
	"podbilling/users/model"
)

type Repository interface {
	CreateUser(login, password string, permissons uint8) error
	GetUserByID(id uint) (model.User, error)
	GetUserByLogin(login string) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id uint) error
	GetAllWorkers() ([]model.User, error)
}
