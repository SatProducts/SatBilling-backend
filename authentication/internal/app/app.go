package app

import (
	handler "podbilling/authentication/internal/http"
	"podbilling/authentication/internal/repository"
	"podbilling/authentication/internal/usecase"
	db "podbilling/authentication/pkg/db"

	"net/http"
	"os"
	"time"
)

type App struct {
	Server  *http.Server
	Handler *handler.AuthHandler
}

func NewApp() *App {
	
	serviceDB := db.ConnectDB(os.Getenv("DB_NAME"))

	repo := repository.NewAuthRepository(
		serviceDB,
	)

	uc := usecase.NewAuthUseCase(
		repo,
		[]byte(os.Getenv("JWT_SIGN_KEY")),
	)

	h := handler.NewAuthHandler(uc)

	server := &http.Server{
		Addr:         "127.0.0.1:8082",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	return &App{
		Server:  server,
		Handler: h,
	}
}
