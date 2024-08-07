package handlers

import (
	"echommerce/internal/models/users_model"
	"echommerce/internal/services"

	"github.com/labstack/echo/v4"

	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) Login(c echo.Context) error {
	var login users_model.UserLogin
	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	user, err := h.UserService.Authenticate(login)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
	}

	token, err := h.UserService.GenerateJWTToken(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "can not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user users_model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	if err := h.UserService.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}
