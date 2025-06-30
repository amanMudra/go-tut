package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Hello, World!")
	app := fiber.New()
	app.Get("/abc", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! 1234")
	})
	app.Listen(":3000")

}
