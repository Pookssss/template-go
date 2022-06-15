package routes

import (
	"template-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.Hello)

	appRoute := app.Group("/api")
	v1App := appRoute.Group("/v1")

	// Users
	users := v1App.Group("/user")
	users.Post("/register", controllers.Register)
	users.Post("/login", controllers.Login)

}
