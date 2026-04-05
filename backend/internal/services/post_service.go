package services

import (
	"errors"
	"strings"

	"rafty/internal/dto"
	"rafty/internal/models"
	"rafty/internal/repositories"
)

type PostService struct {
	PostRepo *repositories.PostRepository
}

func NewPostService(postRepo *repositories.PostRepository) *PostService {
	return &PostService{PostRepo: postRepo}
}

// CreatePost creates a new post with the given details and returns the created post.
func (s *PostService) CreatePost(userID string, req dto.CreatePostRequest) (*models.Post, error) {
	title := normalizeOptionalString(req.Title)
	content := normalizeOptionalString(req.Content)

	if title == nil && content == nil {
		return nil, errors.New("o post precisa ter título ou conteúdo")
	}

	if req.Visibility == "" {
		req.Visibility = "public"
	}

	post := &models.Post{
		AuthorID:   userID,
		Title:      title,
		Content:    content,
		Visibility: req.Visibility,
	}

	err := s.PostRepo.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// normalizeOptionalString trims the input string and returns a pointer to it.
// If the trimmed string is empty, it returns nil.
func normalizeOptionalString(input *string) *string {
	if input == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*input)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}

// GetFeedPosts retrieves a list of posts for the feed, applying pagination and visibility filters.
func (s *PostService) GetFeed(userID string, page, limit int) ([]models.Post, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	return s.PostRepo.GetFeedPosts(userID, limit, offset)
}

// DeletePost performs a soft deletion of a post by its ID, ensuring the user is the author.
func (s *PostService) DeletePost(postID, userID string) error {
	post, err := s.PostRepo.GetPostByID(postID)
	if err != nil {
		return errors.New("post não encontrado")
	}

	if post.AuthorID != userID {
		return errors.New("você não pode deletar este post")
	}

	return s.PostRepo.DeletePost(postID)
}
