package routes

import (
	"rafty/internal/controllers"
	"rafty/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App, postController *controllers.PostController) {
	posts := app.Group("/posts", middlewares.AuthMiddleware())

	posts.Post("/", postController.CreatePost)
	posts.Get("/feed", postController.GetFeed)
	posts.Delete("/:id", postController.DeletePost)
}
