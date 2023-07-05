package routes

import (
	"github.com/abdullah-alaadine/basic-rest-api/database"
	"github.com/abdullah-alaadine/basic-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID      uint     `json:"id"`
	User    *User    `json:"user"`
	Product *Product `json:"product"`
}

func CreateResponseOrder(order *models.Order, user *User, product *Product) Order {
	return Order{
		ID:      order.ID,
		User:    user,
		Product: product,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user models.User
	if err := database.Database.Db.First(&user, order.UserRefer).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err,
		})
	}

	var product models.Product
	if err := database.Database.Db.First(&product, order.ProductRefer).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err,
		})
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(&user)
	responseProduct := CreateResponseProduct(&product)
	responseOrder := CreateResponseOrder(&order, &responseUser, &responseProduct)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order": responseOrder,
	})
}

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order

	database.Database.Db.Find(&orders)

	var responseOrders []Order

	for _, order := range orders {
		var user models.User
		database.Database.Db.First(&user, "id = ?", order.UserRefer)
		var product models.Product
		database.Database.Db.First(&product, "id =?", order.ProductRefer)
		responseUser := CreateResponseUser(&user)
		responseProduct := CreateResponseProduct(&product)
		responseOrder := CreateResponseOrder(&order, &responseUser, &responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": responseOrders,
	})
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var order models.Order
	database.Database.Db.First(&order, id)
	var user models.User
	database.Database.Db.First(&user, "id =?", order.UserRefer)
	var product models.Product
	database.Database.Db.First(&product, "id =?", order.ProductRefer)
	responseUser := CreateResponseUser(&user)
	responseProduct := CreateResponseProduct(&product)
	responseOrder := CreateResponseOrder(&order, &responseUser, &responseProduct)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": responseOrder,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var order models.Order
	database.Database.Db.First(&order, id)

	database.Database.Db.Delete(&order)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
