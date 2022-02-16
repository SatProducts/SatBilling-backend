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

func (repo *Repository) Get(login, password string) (model.User, error) {

	var user model.User

	result := repo.DB.Where(
		"login = ? AND password = ?",
		login, password,
	).First(&user)

	return user, result.Error
}
