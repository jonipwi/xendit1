package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonipwi/xendit-1/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/list", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
