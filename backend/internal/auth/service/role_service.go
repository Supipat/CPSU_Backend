package service

import (
	"cpsu/internal/auth/repository"
)

type RoleService struct {
	RoleRepo *repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{RoleRepo: roleRepo}
}

func (s *RoleService) AssignRole(userID, roleID, assignedBy int) error {
	if err := s.RoleRepo.RemoveRole(userID); err != nil {
		return err
	}

	return s.RoleRepo.AssignRole(userID, roleID, assignedBy)
}
