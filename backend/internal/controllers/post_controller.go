package controllers

import (
	"strconv"

	"rafty/internal/dto"
	"rafty/internal/services"
	"rafty/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	PostService *services.PostService
}

func NewPostController(postService *services.PostService) *PostController {
	return &PostController{PostService: postService}
}

// CreatePostHandler handles the creation of a new post.
func (pc *PostController) CreatePost(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req dto.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Payload inválido")
	}

	post, err := pc.PostService.CreatePost(userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Post criado com sucesso", post)
}

// GetFeedPostsHandler retrieves a list of posts for the feed with pagination and visibility filters.
func (pc *PostController) GetFeed(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	posts, err := pc.PostService.GetFeed(userID, page, limit)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Erro ao buscar feed")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Feed carregado com sucesso", posts)
}

// DeletePostHandler handles the soft deletion of a post by its ID.
func (pc *PostController) DeletePost(c *fiber.Ctx) error {
	postID := c.Params("id")

	err := pc.PostService.DeletePost(postID, c.Locals("user_id").(string))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Post deletado com sucesso", nil)
}
