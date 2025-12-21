package service

import (
	"cpsu/internal/permission/models"
	"cpsu/internal/permission/repository"
	"errors"
)

type PermissionService interface {
	UpdateUserRole(id int, req models.ChangeUserRoleRequest) (*models.UserRole, error)
}

type permissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionService{
		repo: repo,
	}
}

func (s *permissionService) UpdateUserRole(id int, req models.ChangeUserRoleRequest) (*models.UserRole, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}

	if req.RoleID <= 0 {
		return nil, errors.New("invalid role id")
	}

	req.UserID = id
	return s.repo.UpdateUserRole(&req)
}
