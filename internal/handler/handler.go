package handler

import (
	"net/http"
	"pr-reviewer-service/internal/domain"
	"pr-reviewer-service/internal/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		TeamName string `json:"team_name"`
		IsActive *bool  `json:"is_active"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}

	if req.Username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username is required"})
	}
	if req.TeamName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "team_name is required"})
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	user := domain.User{
		Username: req.Username,
		TeamName: req.TeamName,
		IsActive: isActive,
	}

	if err := h.userService.CreateUser(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "user created",
		"user_id": user.UserID,
	})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.userService.GetUser(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	userID := c.Param("id")
	var req domain.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	req.UserID = userID

	if err := h.userService.UpdateUser(c.Request().Context(), req); err != nil {
		if err.Error() == "no rows affected" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user updated"})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	if err := h.userService.DeleteUser(c.Request().Context(), userID); err != nil {
		if err.Error() == "no rows affected" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user deleted"})
}

func (h *UserHandler) ListUsers(c echo.Context) error {
	team := c.QueryParam("team")
	users, err := h.userService.ListUsers(c.Request().Context(), team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
