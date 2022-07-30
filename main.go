package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/onattech/go-postgres-docker-heroku/database"
	"github.com/onattech/go-postgres-docker-heroku/routes"
	_ "github.com/onattech/go-postgres-docker-heroku/swagger"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/products", routes.AllProducts)
	app.Post("/product", routes.AddProduct)
	app.Get("/product/:id", routes.Product)
	app.Put("/product/:id", routes.Update)
	app.Delete("/product/:id", routes.Delete)
	app.Get("/*", swagger.HandlerDefault)
}

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
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${white}${pid} ${cyan}[${time}] ${red}${status} ${latency} ${blue}${method} ${white}${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
	}))

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}
