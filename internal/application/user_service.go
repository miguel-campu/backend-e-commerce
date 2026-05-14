package application

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/tool"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error) {
	existing, _ := s.repo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		ID:           uuid.New(),
		Email:        req.Email,
		PasswordHash: tool.HashSHA256(req.Password),
		FullName:     req.FullName,
		Role:         "user",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := tool.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		User:         *user,
		Token:        token,
		RefreshToken: "dummy-refresh-token",
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req model.CreateUserRequest) (*model.User, error) {
	existing, _ := s.repo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		ID:           uuid.New(),
		Email:        req.Email,
		PasswordHash: tool.HashSHA256(req.Password),
		FullName:     req.FullName,
		Role:         req.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, req model.LoginRequest) (*model.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		// Log for debugging
		println("DEBUG: User not found for email: " + req.Email)
		return nil, errors.New("correo o contraseña inválidos")
	}

	hashedInput := tool.HashSHA256(req.Password)
	if !tool.CheckHashSHA256(req.Password, user.PasswordHash) {
		// Log for debugging
		println("DEBUG: Password mismatch for: " + req.Email)
		println("DEBUG: Input: " + req.Password + " (Hash: " + hashedInput + ")")
		println("DEBUG: DB Hash: " + user.PasswordHash)
		return nil, errors.New("correo o contraseña inválidos")
	}

	token, err := tool.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		User:         *user,
		Token:        token,
		RefreshToken: "dummy-refresh-token",
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	return s.repo.Update(ctx, user)
}

func (s *UserService) UpdatePassword(ctx context.Context, id string, newPassword string) error {
	passwordHash := tool.HashSHA256(newPassword)
	return s.repo.UpdatePassword(ctx, id, passwordHash)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
