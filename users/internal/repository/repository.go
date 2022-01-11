package repository

import (
	"podbilling/users/model"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		DB: db,
	}
}

func (repo *UsersRepository) GetByID(id uint) (model.User, error) {

	var user model.User

	result := repo.DB.First(&user, id)

	return user, result.Error
}

func (repo *UsersRepository) GetByLogin(login string) (model.User, error) {

	var user model.User

	result := repo.DB.Where("login = ?", login).First(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}


func (repo *UsersRepository) Create(login, password string, permissions uint8) error {
	
	user := model.User{
		Login: login,
		Password: password,
		Permissions: permissions,
		Vacation: false,
	}
	
	result := repo.DB.Save(&user)

	return result.Error
}

func (repo *UsersRepository) Update(user model.User) error {
	result := repo.DB.Save(&user)
	return result.Error
}

func (repo *UsersRepository) Delete(id uint) error {
	result := repo.DB.Delete(&model.User{}, id)
	return result.Error
}

func (repo *UsersRepository) GetWorkers() ([]model.User, error) {
	
	var workers []model.User

	result := repo.DB.Where("permissions != ?", model.ADMINISTRATOR).Find(&workers)

	if result.Error != nil {
		return []model.User{}, result.Error
	}

	return workers, nil
}