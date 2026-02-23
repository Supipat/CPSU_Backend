// internal/audit/service/audit_service.go
package service

import (
	"context"

	"cpsu/internal/auth/models"
	"cpsu/internal/auth/repository"
)

type AuditService struct {
	AuditRepo *repository.AuditRepository
}

func NewAuditService(auditRepo *repository.AuditRepository) *AuditService {
	return &AuditService{AuditRepo: auditRepo}
}

func (s *AuditService) GetAllAuditLog(ctx context.Context) ([]models.AuditLogResponse, error) {
	audits, err := s.AuditRepo.GetAllAuditLog(ctx)
	if err != nil {
		return nil, err
	}

	var res []models.AuditLogResponse
	for _, a := range audits {
		res = append(res, models.AuditLogResponse{
			ID:          a.ID,
			UserID:      a.UserID,
			Action:      a.Action,
			Resource:    a.Resource,
			ResourceID:  a.ResourceID,
			Description: Description(a),
			CreatedAt:   a.CreatedAt,
		})
	}

	return res, nil
}

func Description(a models.AuditLog) string {
	switch a.Action {
	case "login":
		return "คุณเข้าสู่ระบบแล้ว"
	case "logout":
		return "คุณออกจากระบบแล้ว"
	case "create":
		return "คุณเพิ่มข้อมูลใหม่แล้ว"
	case "update":
		return "คุณแก้ไขข้อมูลแล้ว"
	case "delete":
		return "คุณลบข้อมูลแล้ว"
	case "assign_role":
		return "คุณให้สิทธิ์ผู้ใช้งานแล้ว"
	default:
		return "มีการดำเนินการในระบบ"
	}
}
