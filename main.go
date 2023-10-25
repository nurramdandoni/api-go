package main

import (
	"api-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Buat instance Fiber
	app := fiber.New()

	// Grupkan rute yang memerlukan otentikasi
	authGroup := app.Group("/auth", middleware.JwtMiddleware)

	// Definisikan route
	authGroup.Get("/cekAuth", func(c *fiber.Ctx) error {
		return c.SendString("Authenticated!")
	})

	// Definisikan route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	// app.Get("/cekAuth", middleware.JwtMiddleware, func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, Fiber!")
	// })

	// Menjalankan server di port 3000
	app.Listen(":3000")

}
