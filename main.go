package main

import (
	"log"

	"github.com/abdullah-alaadine/basic-rest-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", welcome)

	log.Fatal(app.Listen(":8080"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}
