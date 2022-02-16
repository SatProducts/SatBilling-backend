package tasks

import (
	"context"
	"podbilling/model"
)

type UseCase interface {
	Get(ctx context.Context, id uint) (model.Task, error)
	Create(ctx context.Context, task model.Task) error
	Update(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, id uint) error
	GetFor(ctx context.Context, userID uint) ([]model.Task, error)
	GetFrom(ctx context.Context, userID uint) ([]model.Task, error)
}
