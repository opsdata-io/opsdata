package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// AuthenticateJWTAndAPIKey is a middleware function that authenticates a JWT token or API key and sets the user to the context if successful
func AuthenticateJWTAndAPIKey(c *fiber.Ctx) error {
	// Check for JWT token
	authHeader := c.Get("Authorization")
	if authHeader != "" {
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJWTSecret()), nil
		})

		if err == nil && token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT claims"})
			}

			email := claims["email"].(string)
			user, err := utils.GetUserByEmail(email)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
			}

			c.Locals("user", user)
			return c.Next()
		}
	}

	// Check for API key
	accessKey := c.Query("accessKey")
	secretKey := c.Query("secretKey")

	if accessKey == "" {
		accessKey = c.Get("X-API-Key")
	}
	if accessKey != "" && secretKey != "" {
		valid, user, err := utils.VerifyAPIKey(accessKey, secretKey)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify API key"})
		}
		if !valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API key"})
		}

		// Set user in locals or context
		c.Locals("user", user)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid credentials"})
}
