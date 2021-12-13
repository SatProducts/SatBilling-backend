package repository

import (
	"podbilling/authentication/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (repo AuthRepository) GetUser(login string, password uint32) (model.User, error) {

	var user model.User

	result := repo.DB.Where(
		"login = ? AND password = ?",
		login, password,
	).First(&user)

	return user, result.Error
}
