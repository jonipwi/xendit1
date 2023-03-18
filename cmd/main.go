package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonipwi/xendit-1/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	/* app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! 2023 ^o^ ")
	}) */

	app.Listen(":3000")
}
