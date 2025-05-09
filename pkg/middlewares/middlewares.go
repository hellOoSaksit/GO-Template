package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// CheckRole middleware checks if the user has the required role
func CheckRole(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("Template_JWT")
		jwtSecretKey := os.Getenv("JWT_SECRET")

		token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		userRole := claims["Role"].(string)

		if userRole != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Insufficient permissions",
			})
		}

		return c.Next()
	}
}
