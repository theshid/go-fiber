package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/theshid/go-fiber/book"

	"github.com/theshid/go-fiber/database"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect to database!")
	}
	fmt.Println("Database opened successfully!")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}
func main() {
	app := fiber.New()

	setUpRoutes(app)

	app.Listen(":3000")
}
