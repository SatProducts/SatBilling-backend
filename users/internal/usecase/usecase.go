package usecase

import (
	users "podbilling/users/internal"
	"podbilling/users/model"
)

type UsersUseCase struct {
	Repository users.Repository
}

func NewUsersUseCase(repo users.Repository) *UsersUseCase {
	return &UsersUseCase{
		Repository: repo,
	}
}

func (uc *UsersUseCase) GetUser(id uint) (model.User, error) {
	return uc.Repository.GetUserByID(id)
}

func (uc *UsersUseCase) CreateUser(login, password string, permissions uint8) error {

	if login == "" || password == "" {
		return users.EmptyFieldError
	}

	_, err := uc.Repository.GetUserByLogin(login)

	if err == nil {
		return users.UserAlreadyExistsError
	}

	return uc.Repository.CreateUser(
		login,
		password,
		permissions,
	)
}

func (uc *UsersUseCase) UpdateUser(user model.User) error {
	return uc.Repository.UpdateUser(user)
}

func (uc *UsersUseCase) DeleteUser(id uint) error {
	return uc.Repository.DeleteUser(id)
}

func (uc *UsersUseCase) GetAllWorkers() ([]model.User, error) {
	return uc.Repository.GetAllWorkers()
}