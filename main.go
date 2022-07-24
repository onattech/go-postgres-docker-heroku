package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/onattech/go-postgres-docker-heroku/database"
	"github.com/onattech/go-postgres-docker-heroku/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/products", routes.AllProducts)
	app.Post("/product", routes.AddProduct)
	app.Get("/product/:id", routes.Product)
	app.Put("/product/:id", routes.Update)
	app.Delete("/product/:id", routes.Delete)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	//
	port := os.Getenv("PORT")
	log.Println("port", port)

	log.Fatal(app.Listen(":" + "3002"))
}
