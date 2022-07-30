package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onattech/go-postgres-docker-heroku/database"
	"github.com/onattech/go-postgres-docker-heroku/models"
)

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber hellooo!!!")
}

// AddProduct adds a new product
// @Summary     Add a new product
// @Description Register product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product body     models.Product true "Register product"
// @Success     200     {object} models.Product{}
// @Failure     400     {string} string "error"
// @Router      /product [post]
func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&product)

	return c.Status(200).JSON(product)
}

// AllProducts is a function to get all products data from database
// @Summary     Get all products
// @Description Get all products
// @Tags        products
// @Accept      json
// @Produce     json
// @Success     200 {array}  models.Product{}
// @Failure     503     {string} string         "error"
// @Router      /products [get]
func AllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Find(&products)
	return c.Status(200).JSON(products)
}

// Single Product
// Product is a function to get a product by ID
// @Summary     Get product by ID
// @Description Get product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Product ID"
// @Success     200 {object} models.Product{}
// @Failure     404     {object} string         "error"
// @Failure     503 {string} string "error"
// @Router      /product/{id} [get]
func Product(c *fiber.Ctx) error {
	product := models.Product{}
	database.DB.Where("id = ?", c.Params("id")).Find(&product)
	return c.Status(200).JSON(product)
}

// Update Single Product
// Product is a function to update a product by ID
// @Summary     Update product by ID
// @Description Update product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id      path     int            true "Product ID"
// @Param       product body     models.Product true "Update product"
// @Success     200     {string} string         "updated"
// @Failure     404 {object} string "error"
// @Failure     503 {string} string "error"
// @Router      /product/{id} [put]
func Update(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Model(&product).Where("id = ?", c.Params("id")).Updates(product)

	return c.Status(200).JSON("updated")
}

// DeleteProduct function removes a product by ID
// @Summary     Remove product by ID
// @Description Remove product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id  path     int    true "Product ID"
// @Success     200 {string} string "deleted"
// @Failure     404 {string} string "Not Found"
// @Failure     503 {string} string "error"
// @Router      /product/{id} [delete]
func Delete(c *fiber.Ctx) error {
	product := models.Product{}

	database.DB.Where("id = ?", c.Params("id")).Delete(&product)

	return c.Status(200).JSON("deleted")
}
