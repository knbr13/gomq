package routes

import (
	"github.com/abdullah-alaadine/basic-rest-api/database"
	"github.com/abdullah-alaadine/basic-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
}

func CreateResponseProduct(product *models.Product) Product {
	return Product{
		ID:           product.ID,
		Name:         product.Name,
		SerialNumber: product.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := database.Database.Db.Create(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	responseProduct := CreateResponseProduct(&product)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"product": responseProduct,
	})
}
