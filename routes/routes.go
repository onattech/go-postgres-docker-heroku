package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onattech/go-postgres-docker-heroku/database"
	"github.com/onattech/go-postgres-docker-heroku/models"
)

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

// Add Product
func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&product)

	return c.Status(200).JSON(product)
}

// All Products
func AllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Find(&products)
	return c.Status(200).JSON(products)
}

// Single Product
func Product(c *fiber.Ctx) error {
	product := models.Product{}
	database.DB.Where("id = ?", c.Params("id")).Find(&product)
	return c.Status(200).JSON(product)
}

// Update
func Update(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Model(&product).Where("id = ?", c.Params("id")).Updates(product)

	return c.Status(200).JSON("updated")
}

// Delete
func Delete(c *fiber.Ctx) error {
	product := []models.Product{}
	database.DB.Where("id = ?", c.Params("id")).Delete(&product)

	return c.Status(200).JSON("deleted")
}
