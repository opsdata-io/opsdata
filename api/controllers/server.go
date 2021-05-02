package controllers

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mattmattox/opsdata/database"
	"github.com/mattmattox/opsdata/models"
)

func ServerCreate(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var userCheck models.User
	database.DB.Where("id = ?", claims.Issuer).First(&userCheck)
	if userCheck.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User was not found",
		})
	}

	userUuid := userCheck.Uuid
	log.Print(userUuid)

	//Creating the access and secret key pair from Random
	accessKey := RandomKey(8)
	secretKey := RandomKey(32)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	access := models.Key{
		Uuid:      userUuid,
		Accesskey: accessKey,
		Secretkey: secretKey,
	}

	database.DB.Create(&access)

	var accessCheck models.Key

	database.DB.Where("accesskey = ? AND secretKey = ?", accessKey, secretKey).First(&accessCheck)

	if access.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "accesskey failed to be created",
		})
	}

	return c.JSON(access)
}

func ServerVerify(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var key models.Key

	database.DB.Where("accesskey = ? AND secretkey = ?", data["accesskey"], data["secretkey"]).First(&key)

	if key.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Bad key",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
