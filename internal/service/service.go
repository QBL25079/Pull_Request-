package service

import (
	"context"
	"pr-reviewer-service/internal/domain"
	"pr-reviewer-service/internal/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	user.GenerateID()
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, user domain.User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	return s.repo.DeleteUser(ctx, userID)
}

func (s *UserService) ListUsers(ctx context.Context, teamName string) ([]domain.User, error) {
	return s.repo.ListUsers(ctx, teamName)
}
