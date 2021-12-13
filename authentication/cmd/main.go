package main

import (
	"github.com/joho/godotenv"
	"log"

	service "podbilling/authentication/internal/app"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
	
	// Create and run app
	app := service.NewApp()
	app.Route()
}
