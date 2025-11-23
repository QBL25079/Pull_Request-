// internal/service/user_service.go
package service

import (
	"context"
	"pr-reviewer-service/internal/domain"
	"pr-reviewer-service/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, user domain.User) error {
	user.GenerateID()
	return s.userRepo.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, user domain.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	return s.userRepo.DeleteUser(ctx, userID)
}

func (s *UserService) ListUsers(ctx context.Context, teamName string) ([]domain.User, error) {
	return s.userRepo.ListUsers(ctx, teamName)
}
