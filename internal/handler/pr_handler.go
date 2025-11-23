package handler

import (
	"net/http"
	"pr-reviewer-service/internal/domain"
	"pr-reviewer-service/internal/service"

	"github.com/labstack/echo/v4"
)

type PRHandler struct {
	prService *service.PRService
}

func NewPRHandler(prService *service.PRService) *PRHandler {
	return &PRHandler{prService: prService}
}

func (h *PRHandler) CreatePR(c echo.Context) error {
	var req struct {
		PullRequestName string  `json:"pull_request_name"`
		AuthorID        string  `json:"author_id"`
		Reviewer1ID     *string `json:"reviewer1_id,omitempty"`
		Reviewer2ID     *string `json:"reviewer2_id,omitempty"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}

	if req.PullRequestName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "pull_request_name is required"})
	}
	if req.AuthorID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "author_id is required"})
	}

	pr := domain.PullRequest{
		PullRequestName: req.PullRequestName,
		AuthorID:        req.AuthorID,
		Reviewer1ID:     req.Reviewer1ID,
		Reviewer2ID:     req.Reviewer2ID,
	}

	if err := h.prService.CreatePullRequest(c.Request().Context(), pr); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message":            "pull request created",
		"pull_request_id":    pr.PullRequestID,
		"status":             pr.Status,
		"assigned_reviewers": pr.AssignedReviewers,
	})
}
