package main

import (
	"log"
	"os"

	"rafty/internal/config"
	"rafty/internal/controllers"
	"rafty/internal/repositories"
	"rafty/internal/routes"
	"rafty/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: .env não encontrado, usando variáveis do sistema")
	}

	db := config.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)
	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	routes.SetupAuthRoutes(app, authController)
	routes.SetupPostRoutes(app, postController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor rodando na porta " + port)
	log.Fatal(app.Listen(":" + port))
}
