package middleware

import (
	"ma-backend-training/internal/service"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware is a middleware function for validating JWT tokens
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "missing token",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token format",
		})
	}

	username, err := service.ParseJWT(tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	// Store username in locals for use in handlers
	c.Locals("username", username)

	return c.Next()
}
