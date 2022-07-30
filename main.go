package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onattech/go-postgres-docker-heroku/database"
	"github.com/onattech/go-postgres-docker-heroku/routes"
	_ "github.com/onattech/go-postgres-docker-heroku/swagger"
)

// @title       Products App
// @version     1.0
// @description This is a simple API for Products

// @contact.name Onat

// @license.name MIT
// @license.url  https://mit-license.org/

// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	database.ConnectDb()

	app := routes.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
