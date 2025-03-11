package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-industry-project/internal/user"
)

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var newUser user.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	hashedPwd, err := HashPassword(newUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not hash password"})
	}
	newUser.Password = hashedPwd

	if err := h.AuthService.UserRepo.Create(&newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User already exists"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	user, err := h.AuthService.UserRepo.FindByEmail(credentials.Email)
	if err != nil || !CheckPassword(user.Password, credentials.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := GenerateToken(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
