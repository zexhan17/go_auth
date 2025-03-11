package db

import (
	"log"
	"os"

	"github.com/zexhan17/go_auth/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run Migrations
	err = database.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	DB = database
	log.Println("Connected to Neon PostgreSQL and migrations applied!")
}
