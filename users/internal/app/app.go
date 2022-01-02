package app

import (
	handler "podbilling/users/internal/http"
	"podbilling/users/internal/repository"
	"podbilling/users/internal/usecase"
	db "podbilling/users/pkg/db"
	"podbilling/users/model"

	"net/http"
	"os"
	"time"
)

type App struct {
	JwtSignKey []byte
	Server  *http.Server
	Handler *handler.UsersHandler
}

func NewApp() *App {
	
	serviceDB := db.ConnectDB()
	serviceDB.AutoMigrate(&model.User{})

	repo := repository.NewUsersRepository(
		serviceDB,
	)

	uc := usecase.NewUsersUseCase(
		repo,
	)

	h := handler.NewUsersHandler(
		uc,
	)

	server := &http.Server{
		Addr:         os.Getenv("HOST"),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	return &App{
		JwtSignKey: []byte(os.Getenv("JWT_SIGN_KEY")),
		Server:  server,
		Handler: h,
	}
}
