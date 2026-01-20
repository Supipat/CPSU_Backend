package models

import "time"

type User struct {
	UserID       int        `json:"user_id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	LastLogin    *time.Time `json:"last_login"`
}

type UserResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
	Name     string `json:"name"`
}

type UserQueryParam struct {
	Search string `form:"search"`
	Limit  int    `form:"limit"`
	UserID int    `form:"user_id"`
	RoleID int    `json:"role_id"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
}
