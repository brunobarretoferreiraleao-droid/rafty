package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"rafty/internal/dto"
	"rafty/internal/models"
	"rafty/internal/repositories"
)

type PostService struct {
	PostRepo *repositories.PostRepository
}

func (s *PostService) CreatePost(ctx context.Context, userID string, req dto.CreatePostRequest) (*models.Post, error) {
	if req.Content == nil && req.Title == nil {
		return nil, fmt.Errorf("post precisa de conteúdo ou título")
	}

	if req.Visibility == "" {
		req.Visibility = "public"
	}

	post := &models.Post{
		AuthorID:   userID,
		Title:      req.Title,
		Content:    req.Content,
		Visibility: req.Visibility,
	}

	err := s.PostRepo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) DeletePost(ctx context.Context, postID string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := s.PostRepo.DeletePost(ctx, postID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("post não encontrado ou não pertence ao usuário")
		}
		return err
	}

	return nil
}

func NewPostService(postRepo *repositories.PostRepository) *PostService {
	return &PostService{PostRepo: postRepo}
}

func (s *PostService) GetFeed(ctx context.Context, userID string, page, limit int) ([]models.Post, error) {
	// 🧠 Proteções básicas
	if page < 1 {
		page = 1
	}

	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	// ⏱️ Timeout (evita request travado)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return s.PostRepo.GetFeedPosts(ctx, userID, limit, offset)
}
