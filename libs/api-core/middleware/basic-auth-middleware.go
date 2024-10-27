package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func BasicAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Remove Bearer prefix
		tokenString := token[7:]

		fmt.Println(tokenString, "<< iki token string")

		// // Parse and validate JWT token
		// _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	// Replace with your actual secret key
		// 	return []byte("your-secret-key"), nil
		// })

		// if err != nil {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": "Invalid token",
		// 	})
		// }

		return c.Next()
	}
}
