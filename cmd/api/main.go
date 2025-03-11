package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/zexhan17/go_auth/internal/auth"
	"github.com/zexhan17/go_auth/internal/db"
	"github.com/zexhan17/go_auth/internal/middleware"
	"github.com/zexhan17/go_auth/internal/user"
)

func main() {
	godotenv.Load()
	db.ConnectDB()

	userRepo := user.NewUserRepository(db.DB)
	authService := auth.NewAuthService(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	e := echo.New()

	// Public Routes
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// Protected Route
	e.GET("/profile", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome to your profile!"})
	}, middleware.JWTMiddleware)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
