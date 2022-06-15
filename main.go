package main

import (
	"log"
	"template-go/database"
	"template-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":5000"))
}
