package tasks

import (
	"podbilling/model"
)

type Repository interface {
	Get(id uint) (model.Task, error)
	Create(task model.Task) error
	Update(task model.Task) error
	Delete(id uint) error
	GetFor(userID uint) ([]model.Task, error)
	GetFrom(userID uint) ([]model.Task, error)
	GetMinimal() (uint, error)
}
