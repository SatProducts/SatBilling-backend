package repository

import (
	"podbilling/model"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repo *Repository) Get(id uint) (model.Task, error) {
	var task model.Task
	result := repo.DB.First(&task, id)
	return task, result.Error
}

func (repo *Repository) Create(task model.Task) error {
	result := repo.DB.Create(&task)
	return result.Error
}

func (repo *Repository) Update(task model.Task) error {
	result := repo.DB.Save(&task)
	return result.Error
}

func (repo *Repository) Delete(id uint) error {
	result := repo.DB.Delete(&model.Task{}, id)
	return result.Error
}

func (repo *Repository) GetFor(userID uint) ([]model.Task, error) {

	var tasks []model.Task

	result := repo.DB.Where(
		"for_user = ?",
		userID,
	).Find(&tasks)

	return tasks, result.Error
}

func (repo *Repository) GetFrom(userID uint) ([]model.Task, error) {

	var tasks []model.Task

	result := repo.DB.Where(
		"from_user = ?",
		userID,
	).Find(&tasks)

	return tasks, result.Error
}

func (repo *Repository) GetMinimal() (uint, error) {

	var user model.User

	result := repo.DB.Where(
		"permissions = 1 AND tasks = (SELECT MIN(tasks) FROM users)",
	).First(&user)

	return user.ID, result.Error
}
