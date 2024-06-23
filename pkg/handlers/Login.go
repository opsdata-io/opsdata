package handlers

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/pkg/models"
	"github.com/opsdata-io/opsdata/pkg/utils"

	_ "github.com/swaggo/fiber-swagger"
)

// Login handles user authentication and returns a JWT token.
// @Summary User login
// @Description Logs in a user and returns a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]interface{} "Successful login"
// @Failure 400 {object} map[string]interface{} "Invalid request format"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/login [post]
func Login(c *fiber.Ctx) error {
	// Parse the request body into a User struct
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Cannot parse JSON"})
	}

	// Verify user credentials against stored data
	if !utils.VerifyCredentials(user.Email, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{"error": "Invalid credentials"})
	}

	// Generate JWT token with email claim and expiry time of 72 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the JWT token with a secret key and handle any errors
	tokenString, err := token.SignedString([]byte(utils.GetJWTSecret()))
	if err != nil {
		log.Println("Error generating token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Could not generate token"})
	}

	// Return the JWT token as JSON response
	return c.JSON(map[string]interface{}{"token": tokenString})
}
