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

func (repo *Repository) GetByID(id uint) (model.User, error) {

	var user model.User

	result := repo.DB.First(&user, id)

	return user, result.Error
}

func (repo *Repository) GetByLogin(login string) (model.User, error) {

	var user model.User

	result := repo.DB.Where("login = ?", login).First(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (repo *Repository) GetWorkers() ([]model.User, error) {

	var workers []model.User

	result := repo.DB.Where("permissions != ?", model.ADMINISTRATOR).Find(&workers)

	if result.Error != nil {
		return []model.User{}, result.Error
	}

	return workers, nil
}

func (repo *Repository) Create(user model.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}

func (repo *Repository) Update(user model.User) error {
	result := repo.DB.Save(&user)
	return result.Error
}

func (repo *Repository) Delete(id uint) error {
	result := repo.DB.Delete(&model.User{}, id)
	return result.Error
}
