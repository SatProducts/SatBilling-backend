package main

import (
	service "podbilling/users/internal/app"
)

func main() {
	// Create and run app
	app := service.NewApp()
	app.Route()
}
