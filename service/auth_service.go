package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"role-based-access-control-service/db"
	"role-based-access-control-service/dto"
	"role-based-access-control-service/models"
	"role-based-access-control-service/pkg/auth"
	"role-based-access-control-service/pkg/apperrors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
}

type AuthService struct {
	repo      AuthRepository
	jwtSecret string
}

func NewAuthService(repo AuthRepository, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthData, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Name = strings.TrimSpace(req.Name)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var name *string
	if req.Name != "" {
		name = &req.Name
	}

	user := &models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     name,
		Role:     "user",
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		if db.IsDuplicateError(err) {
			return nil, apperrors.ErrEmailAlreadyExists
		}
		return nil, err
	}

	token, err := auth.GenerateToken(s.jwtSecret, user.ID, user.Email, user.Role, time.Hour)
	if err != nil {
		return nil, err
	}

	return &dto.AuthData{User: dto.UserFromModel(user), Token: token}, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthData, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(s.jwtSecret, user.ID, user.Email, user.Role, time.Hour)
	if err != nil {
		return nil, err
	}

	return &dto.AuthData{User: dto.UserFromModel(user), Token: token}, nil
}

func (s *AuthService) Me(ctx context.Context, userID int) (*dto.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUnauthorized
		}
		return nil, err
	}

	resp := dto.UserFromModel(user)
	return &resp, nil
}
