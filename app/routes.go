package app

import (
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func (app *App) Run() {
	// r.Schemes("https")

	app.RouteUsers()
	app.RouteAuth()
	app.RouteTasks()

	app.Server.Get("/", func(c *fiber.Ctx) error {
		return c.Render("greet", nil)
	})

	log.Fatal(
		app.Server.Listen(":8080"),
	)
}

func (app *App) RouteUsers() {

	sub := app.Server.Group("/user")

	sub.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SIGN_KEY")),
	}))

	sub.Get("/me/", app.UsersHandler.GetSelf)
	sub.Get("/:id/", app.UsersHandler.Get)
	sub.Get("/workers/", app.UsersHandler.GetWorkers)
	sub.Post("/", app.UsersHandler.Create)
	sub.Put("/:id/", app.UsersHandler.Update)
	sub.Delete("/:id/", app.UsersHandler.Delete)
}

func (app *App) RouteAuth() {

	sub := app.Server.Group("/auth")

	sub.Post("/login/", app.AuthHandler.Login)
}

func (app *App) RouteTasks() {

	sub := app.Server.Group("/task")

	sub.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SIGN_KEY")),
	}))

	sub.Get("/:id/", app.TasksHandler.Get)
	sub.Post("/", app.TasksHandler.Create)
	sub.Put("/:id/", app.TasksHandler.Update)
	sub.Delete("/:id/", app.TasksHandler.Delete)
	sub.Get("/for/", app.TasksHandler.GetFor)
	sub.Get("/from/", app.TasksHandler.GetFrom)
}
