package app

import (
	"os"
	authhttp "podbilling/internal/authentication/http"
	authrepo "podbilling/internal/authentication/repository"
	authuc "podbilling/internal/authentication/usecase"

	usershttp "podbilling/internal/users/http"
	usersrepo "podbilling/internal/users/repository"
	usersuc "podbilling/internal/users/usecase"

	taskshttp "podbilling/internal/tasks/http"
	tasksrepo "podbilling/internal/tasks/repository"
	tasksuc "podbilling/internal/tasks/usecase"

	"gorm.io/gorm"
)

func RegisterAuthService(db *gorm.DB) *authhttp.Handler {

	repo := authrepo.NewRepository(
		db,
	)

	uc := authuc.NewUseCase(
		repo,
		[]byte(os.Getenv("JWT_SIGN_KEY")),
	)

	return authhttp.NewHandler(uc)
}

func RegisterUsersService(db *gorm.DB) *usershttp.Handler {

	repo := usersrepo.NewRepository(
		db,
	)

	uc := usersuc.NewUseCase(
		repo,
	)

	return usershttp.NewHandler(uc)
}

func RegisterTasksService(db *gorm.DB) *taskshttp.Handler {

	repo := tasksrepo.NewRepository(
		db,
	)

	uc := tasksuc.NewUseCase(
		repo,
	)

	return taskshttp.NewHandler(uc)
}
