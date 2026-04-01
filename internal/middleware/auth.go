package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("JWT_SECRET_KEY is not set")
	}
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: jwtErrorHandler,
	})
}

func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken, ok := c.Locals("user").(*jwt.Token)
		if !ok || userToken == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access",
			})
		}

		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access",
			})
		}

		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: insufficient permissions",
			})
		}

		return c.Next()
	}
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized access",
	})
}
