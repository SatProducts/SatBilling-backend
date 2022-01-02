package main

import (
	service "podbilling/authentication/internal/app"
)

func main() {
	// Create and run app
	app := service.NewApp()
	app.Route()
}
