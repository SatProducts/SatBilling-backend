package main

import (
	"log"
	"podbilling/app"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file does not exists")
	}

	server := app.NewApp()
	server.Run()
}
