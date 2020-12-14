package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("A Single Books")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Books")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete a Book")
}
