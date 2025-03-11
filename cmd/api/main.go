package main

import (
	"os"

	"https://github.com/zexhan17/go_auth/internal/auth"
	"https://github.com/zexhan17/go_auth/internal/db"
	"https://github.com/zexhan17/go_auth/internal/user"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	db.ConnectDB()

	userRepo := user.NewUserRepository(db.DB)
	authService := auth.NewAuthService(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	e := echo.New()
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
