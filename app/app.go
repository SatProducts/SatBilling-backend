package app

import (
	authhttp "podbilling/internal/authentication/http"
	taskshttp "podbilling/internal/tasks/http"
	usershttp "podbilling/internal/users/http"
	"podbilling/model"
	"podbilling/pkg/db"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type App struct {
	Server *fiber.App

	UsersHandler *usershttp.Handler
	AuthHandler  *authhttp.Handler
	TasksHandler *taskshttp.Handler
}

func NewApp() *App {

	appDB := db.InitDB()

	appDB.AutoMigrate(&model.User{})
	appDB.AutoMigrate(&model.Task{})

	return &App{
		Server: fiber.New(fiber.Config{
			Views: html.New("./web", ".html"),
		}),
		TasksHandler: RegisterTasksService(appDB),
		UsersHandler: RegisterUsersService(appDB),
		AuthHandler:  RegisterAuthService(appDB),
	}
}
