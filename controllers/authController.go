package controllers

import (
	"template-go/database"
	"template-go/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var SecreteKey = []byte("secret")

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CheckAuthentication(c *fiber.Ctx) error {

	cookie := c.Cookies("token")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecreteKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": 401, "message": "Unauthorized"})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.SendString("You are authenticated")
}
