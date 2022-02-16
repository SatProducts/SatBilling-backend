package users

import (
	"context"
	"podbilling/model"
)

type UseCase interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id uint) (model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id uint) error
	GetWorkers(ctx context.Context) ([]model.User, error)
}
