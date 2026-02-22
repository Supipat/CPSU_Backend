package service

import (
	"cpsu/internal/auth/repository"
	"strconv"
)

type RoleService struct {
	RoleRepo  *repository.RoleRepository
	AuditRepo *repository.AuditRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{RoleRepo: roleRepo}
}

func (s *RoleService) AssignRole(userID, roleID, assignedBy int, ipAddress string, userAgent string) error {
	if err := s.RoleRepo.RemoveRole(userID); err != nil {
		return err
	}

	if err := s.RoleRepo.AssignRole(userID, roleID, assignedBy); err != nil {
		return err
	}

	_ = s.AuditRepo.LogAudit(
		assignedBy, "assign_role", "user", strconv.Itoa(userID),
		map[string]interface{}{
			"user_id": userID,
			"role_id": roleID,
		},
		ipAddress,
		userAgent,
	)

	return nil
}
