package usecase

import (
	"context"
	"podbilling/internal/tasks"
	"podbilling/model"
)

type UseCase struct {
	Repository tasks.Repository
}

func NewUseCase(repo tasks.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

func (uc *UseCase) Get(ctx context.Context, id uint) (model.Task, error) {
	return uc.Repository.Get(id)
}

func (uc *UseCase) Create(ctx context.Context, task model.Task) error {

	if task.Title == "" || task.Address == "" || task.Text == "" {
		return tasks.EmptyFieldError
	}

	forID, err := uc.Repository.GetMinimal()

	if err != nil {
		return err
	}

	task.ForUser = forID

	return uc.Repository.Create(task)
}

func (uc *UseCase) Update(ctx context.Context, task model.Task) error {

	if task.Title == "" || task.Address == "" || task.Text == "" {
		return tasks.EmptyFieldError
	}

	return uc.Repository.Update(task)
}

func (uc *UseCase) Delete(ctx context.Context, id uint) error {
	return uc.Repository.Delete(id)
}

func (uc *UseCase) GetFor(ctx context.Context, userID uint) ([]model.Task, error) {
	return uc.Repository.GetFor(userID)
}

func (uc *UseCase) GetFrom(ctx context.Context, userID uint) ([]model.Task, error) {
	return uc.Repository.GetFrom(userID)
}
