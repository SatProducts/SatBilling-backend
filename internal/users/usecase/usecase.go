package usecase

import (
	"context"
	users "podbilling/internal/users"
	"podbilling/model"
)

type UseCase struct {
	Repository users.Repository
}

func NewUseCase(repo users.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

func (uc *UseCase) Get(ctx context.Context, id uint) (model.User, error) {
	return uc.Repository.GetByID(id)
}

func (uc *UseCase) Create(ctx context.Context, user model.User) error {

	if user.Login == "" || user.Password == "" {
		return users.EmptyFieldError
	}

	_, err := uc.Repository.GetByLogin(user.Login)

	if err == nil {
		return users.UserAlreadyExistsError
	}

	return uc.Repository.Create(user)
}

func (uc *UseCase) Update(ctx context.Context, user model.User) error {
	return uc.Repository.Update(user)
}

func (uc *UseCase) Delete(ctx context.Context, id uint) error {
	return uc.Repository.Delete(id)
}

func (uc *UseCase) GetWorkers(ctx context.Context) ([]model.User, error) {
	return uc.Repository.GetWorkers()
}
