package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/onattech/go-postgres-docker-heroku/handlers"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${white}${pid} ${cyan}[${time}] ${red}${status} ${latency} ${blue}${method} ${white}${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
	}))

	api := app.Group("/api")
	api.Get("/hello", handlers.Hello)
	api.Get("/products", handlers.AllProducts)
	api.Post("/product", handlers.AddProduct)
	api.Get("/product/:id", handlers.Product)
	api.Put("/product/:id", handlers.Update)
	api.Delete("/product/:id", handlers.Delete)
	app.Get("/*", swagger.HandlerDefault)

	return app
}
