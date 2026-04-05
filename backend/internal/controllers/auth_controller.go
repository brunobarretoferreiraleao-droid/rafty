package controllers

import (
	"rafty/internal/models"
	"rafty/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Handler for user registration
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req models.RegisterRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	resp, err := c.AuthService.Register(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

// Handler for user login
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req models.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	resp, err := c.AuthService.Login(req)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(resp)
}

// Me handler to get current user info
func (c *AuthController) Me(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)

	user, err := c.AuthService.UserRepo.GetByID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return ctx.JSON(user)
}
