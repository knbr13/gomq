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

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if err := database.Database.Db.Find(&products); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	var productsResponse []Product

	for _, product := range products {
		productsResponse = append(productsResponse, CreateResponseProduct(&product))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": productsResponse,
	})
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var product models.Product
	if err := database.Database.Db.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productResponse := CreateResponseProduct(&product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": productResponse,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var product models.Product
	if err := database.Database.Db.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serialNumber"`
	}

	var updateProduct UpdateProduct

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	if err := database.Database.Db.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	responseProduct := CreateResponseProduct(&product)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": responseProduct,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var product models.Product
	if err := database.Database.Db.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
