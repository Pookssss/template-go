package controllers

import (
	"strconv"
	"template-go/database"
	"template-go/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterEndpoint(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	user := models.User{
		Username:  data["username"],
		Email:     data["email"],
		Password:  password,
		CreatedAt: time.Now(),
	}

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id != 0 {
		return c.JSON(fiber.Map{"status": 400, "message": "Email already exists"})
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{"status": 200, "data": user})
}

func LoginEndpoint(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"status": 404, "message": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": 401, "message": "Incorect password"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecreteKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"status": 500, "message": "Internal server error"})
	}

	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(cookie)

	data = map[string]string{"token": token}

	return c.JSON(fiber.Map{"status": 200, "data": data, "message": "success"})
}

func UserEndpoint(c *fiber.Ctx) error {

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

	return c.JSON(fiber.Map{"status": 200, "data": user})

}

func LogoutEndpoint(c *fiber.Ctx) error {
	// cookie := c.Cookies("token")
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"status": 200, "message": "success"})
}
