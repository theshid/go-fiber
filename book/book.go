package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("A Single Books")
}

func NewBooks(c *fiber.Ctx) error {
	return c.SendString("New Books")
}

func DeleteBooks(c *fiber.Ctx) error {
	return c.SendString("Delete a Book")
}
