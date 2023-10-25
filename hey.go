package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Buat instance Fiber
	app := fiber.New()

	// Definisikan route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Menjalankan server di port 3000
	app.Listen(":3000")

}
