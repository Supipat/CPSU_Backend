package models

type UserRole struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	Name     string `json:"name"`
}

type ChangeUserRoleRequest struct {
	UserID int `json:"-"`
	RoleID int `json:"role_id" binding:"required"`
}
