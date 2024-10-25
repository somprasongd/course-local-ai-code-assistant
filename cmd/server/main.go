package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a simple route that responds with "Hello World!"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server
	fmt.Println("Server is running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
