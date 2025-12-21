package repository

import (
	"database/sql"

	"cpsu/internal/permission/models"
)

type PermissionRepository interface {
	UpdateUserRole(req *models.ChangeUserRoleRequest) (*models.UserRole, error)
}

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) UpdateUserRole(req *models.ChangeUserRoleRequest) (*models.UserRole, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM user_roles WHERE user_id = $1`, req.UserID)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		INSERT INTO user_roles (user_id, role_id)
		VALUES ($1, $2)
	`, req.UserID, req.RoleID)
	if err != nil {
		return nil, err
	}

	var res models.UserRole
	err = tx.QueryRow(`
		SELECT 
			u.user_id, u.username, r.role_id, r.name
		FROM users u
		JOIN user_roles ur ON u.user_id = ur.user_id
		JOIN roles r ON ur.role_id = r.role_id
		WHERE u.user_id = $1
	`, req.UserID).Scan(
		&res.UserID, &res.Username, &res.RoleID, &res.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &res, tx.Commit()
}
