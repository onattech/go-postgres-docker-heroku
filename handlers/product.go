package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onattech/go-postgres-docker-heroku/database"
	"github.com/onattech/go-postgres-docker-heroku/models"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber hellooo!!!")
}

// AddProduct adds a new product
// @Summary     Add a new product
// @Description Add a new product
// @Tags        Products
// @Accept      json
// @Produce     json
// @Param       product body     models.Product true "Register product"
// @Success     200     {object} models.Product{}
// @Failure     400     {string} string
// @Router      /api/product [post]
func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(product)

	return c.Status(200).JSON(product)
}

// GetAllProducts is a function to get all products data from database
// @Summary     Get all products
// @Description Get all products
// @Tags        Products
// @Accept      json
// @Produce     json
// @Success     200 {array}  models.Product{}
// @Failure     503     {string} string
// @Router      /api/products [get]
func GetAllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Find(&products)
	return c.Status(200).JSON(products)
}

// Single GetSingleProduct
// GetSingleProduct is a function to get a product by ID
// @Summary     Get product by ID
// @Description Get product by ID
// @Tags        Products
// @Accept      json
// @Produce     json
// @Param       id  path     int true "GetSingleProduct ID"
// @Success     200 {object} models.Product{}
// @Failure     404     {object} string
// @Failure     503 {string} string
// @Router      /api/product/{id} [get]
func GetSingleProduct(c *fiber.Ctx) error {
	product := models.Product{}
	database.DB.Where("id = ?", c.Params("id")).Find(&product)
	return c.Status(200).JSON(product)
}

// UpdateProduct Single Product
// Product is a function to update a product by ID
// @Summary     UpdateProduct product by ID
// @Description UpdateProduct product by ID
// @Tags        Products
// @Accept      json
// @Produce     json
// @Param       id      path     int            true "Product ID"
// @Param       product body     models.Product true "UpdateProduct product"
// @Success     200     {string} string
// @Failure     404 {object} string
// @Failure     503 {string} string
// @Router      /api/product/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
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
// @Tags        Products
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Product ID"
// @Success     200 {string} string
// @Failure     404 {string} string
// @Failure     503 {string} string
// @Router      /api/product/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	product := models.Product{}

	database.DB.Where("id = ?", c.Params("id")).Delete(&product)

	return c.Status(200).JSON("deleted")
}
