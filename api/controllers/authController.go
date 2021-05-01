package controllers

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mattmattox/opsdata/database"
	"github.com/mattmattox/opsdata/models"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomKey(length int) string {
	return StringWithCharset(length, charset)
}

func CreateApiKey(c *fiber.Ctx) error {
	var data map[string]string

	//Need to add section to verify user is logged into and get their Uuid

	//Creating the access and secret key pair from Random
	accessKey := RandomKey(8)
	secretKey := RandomKey(32)

	UserUuid := data["UserUuid"]
	var userCheck models.User
	database.DB.Where("uuid = ?", data["UserUuid"]).First(&userCheck)
	if userCheck.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User UUID was not found",
		})
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	secretHash, err := bcrypt.GenerateFromPassword([]byte(secretKey), 14)
	if err != nil {
		return err
	}

	uuidWithHyphen := uuid.New()
	uuid := uuidWithHyphen.String()

	access := models.Access{
		Uuid:      uuid,
		UserUuid:  UserUuid,
		Accesskey: accessKey,
		Secretkey: secretHash,
	}

	database.DB.Create(&access)

	var accessCheck models.Access

	database.DB.Where("accesskey = ?", data["accesskey"]).First(&accessCheck)

	if access.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "accesskey failed to be created",
		})
	}

	return c.JSON(access)
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return err
	}

	uuidWithHyphen := uuid.New()
	uuid := uuidWithHyphen.String()

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Uuid:     uuid,
		Password: password,
	}

	database.DB.Create(&user)

	var userCheck models.User

	database.DB.Where("email = ?", data["email"]).First(&userCheck)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "email already registered",
		})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
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

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
