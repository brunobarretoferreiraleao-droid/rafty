package services

import (
	"errors"
	"strings"

	"rafty/internal/models"
	"rafty/internal/repositories"
	"rafty/internal/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

// Method to register a new user
func (s *AuthService) Register(req models.RegisterRequest) (*models.AuthResponse, error) {
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if len(req.Username) < 3 {
		return nil, errors.New("username must have at least 3 characters")
	}

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if len(req.Password) < 6 {
		return nil, errors.New("password must have at least 6 characters")
	}

	if _, err := s.UserRepo.GetByUsername(req.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	if _, err := s.UserRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	if err := s.UserRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

// Method do login an existing user
func (s *AuthService) Login(req models.LoginRequest) (*models.AuthResponse, error) {
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(req.Password, user.PasswordHash); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}
