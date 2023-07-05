package main

import (
	"log"

	"github.com/abdullah-alaadine/basic-rest-api/database"
	"github.com/abdullah-alaadine/basic-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	// User routes
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}
