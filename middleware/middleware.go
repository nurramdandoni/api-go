package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware JWT
func JwtMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if !t.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
