package main

import (
	"log"
	"os"
	"template-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
	}))

	routes.SetupRoutes(app)
	MODE := os.Getenv("MODE")

	if MODE == "production" {
		log.Fatal(app.Listen(":5000"))
	} else {
		log.Fatal(app.Listen("127.0.0.1:5000"))
	}

}
