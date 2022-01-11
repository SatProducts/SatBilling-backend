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

func (uc *UsersUseCase) Get(id uint) (model.User, error) {
	return uc.Repository.GetByID(id)
}

func (uc *UsersUseCase) Create(login, password string, permissions uint8) error {

	if login == "" || password == "" {
		return users.EmptyFieldError
	}

	_, err := uc.Repository.GetByLogin(login)

	if err == nil {
		return users.UserAlreadyExistsError
	}

	return uc.Repository.Create(
		login,
		password,
		permissions,
	)
}

func (uc *UsersUseCase) Update(user model.User) error {
	return uc.Repository.Update(user)
}

func (uc *UsersUseCase) Delete(id uint) error {
	return uc.Repository.Delete(id)
}

func (uc *UsersUseCase) GetWorkers() ([]model.User, error) {
	return uc.Repository.GetWorkers()
}