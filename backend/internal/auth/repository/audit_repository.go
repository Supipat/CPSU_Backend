package repository

import (
	"database/sql"
	"encoding/json"
)

type AuditRepository struct {
	db *sql.DB
}

func NewAuditRepository(db *sql.DB) *AuditRepository {
	return &AuditRepository{db: db}
}

func (r *AuditRepository) LogAudit(userID int, action string, resource string, resourceID string, details map[string]interface{}, ipAddress string, userAgent string) error {
	var detailsJSON []byte
	if details != nil {
		detailsJSON, _ = json.Marshal(details)
	}

	query := `
		INSERT INTO audit_logs(user_id, action, resource, resource_id, details, ip_address, user_agent)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(
		query, userID, action, resource, resourceID,
		detailsJSON, ipAddress, userAgent,
	)

	return err
}
