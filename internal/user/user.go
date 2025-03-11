package user

import "gorm.io/gorm"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`
	Role     Role   `gorm:"default:user" json:"role"`
}
