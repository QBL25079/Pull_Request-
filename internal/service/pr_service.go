// internal/service/pr_service.go
package service

import (
	"context"
	"pr-reviewer-service/internal/domain"
	"pr-reviewer-service/internal/repository"
	"time"
)

type PRService struct {
	repo repository.Repository
}

func NewPRService(repo repository.Repository) *PRService {
	return &PRService{repo: repo}
}

func (s *PRService) CreatePullRequest(ctx context.Context, pr domain.PullRequest) error {
	pr.GenerateID()

	pr.Status = domain.PRStatusOpen
	now := time.Now()
	pr.CreatedAt = &now

	pr.AssignedReviewers = []string{}
	if pr.Reviewer1ID != nil && *pr.Reviewer1ID != "" {
		pr.AssignedReviewers = append(pr.AssignedReviewers, *pr.Reviewer1ID)
	}
	if pr.Reviewer2ID != nil && *pr.Reviewer2ID != "" {
		pr.AssignedReviewers = append(pr.AssignedReviewers, *pr.Reviewer2ID)
	}

	return s.repo.CreatePullRequest(ctx, pr)
}
