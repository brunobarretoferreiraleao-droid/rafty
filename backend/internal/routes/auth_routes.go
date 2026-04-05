package routes

import (
	"rafty/internal/controllers"
	"rafty/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, authController *controllers.AuthController) {
	auth := app.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	auth.Get("/me", middlewares.AuthMiddleware(), authController.Me)
}
