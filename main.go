package main

import (
	"github.com/abdullah-alaadine/basic-rest-api/database"
	"github.com/abdullah-alaadine/basic-rest-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":8000")
}
