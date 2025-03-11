package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zexhan17/go_auth/internal/user"
)

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var newUser user.User

	// Debug: Print the request body
	fmt.Println("Received request body")

	// Try to parse JSON request
	if err := c.Bind(&newUser); err != nil {
		fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Debug: Print parsed user
	fmt.Printf("Parsed user: %+v\n", newUser)

	// Ensure fields are not empty
	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}

	// Hash the password before saving
	hashedPwd, err := HashPassword(newUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not hash password"})
	}
	newUser.Password = hashedPwd

	// Save user to the database
	if err := h.AuthService.UserRepo.Create(&newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User already exists"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}
