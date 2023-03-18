package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonipwi/xendit-1/database"
	"github.com/jonipwi/xendit-1/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Jonny :)!")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	fact := []models.Fact{}

	database.DB.Db.Find(&fact)

	return c.Status(200).JSON(fact)
}
