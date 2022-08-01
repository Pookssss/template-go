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
	users := v1App.Group("")
	users.Post("/register", controllers.RegisterEndpoint)
	users.Post("/login", controllers.LoginEndpoint)
	users.Get("/user", controllers.UserEndpoint)
	users.Post("/logout", controllers.LogoutEndpoint)

	agora := v1App.Group("/agora")
	agora.Get("/projects/:customerKey/:customerSecret", controllers.ListProjectEndpoint)
	agora.Get("/rules/:appId/:customerKey/:customerSecret", controllers.RuleListByAppIDEndpoint)

}
