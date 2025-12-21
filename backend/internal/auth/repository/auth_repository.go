package repository

import (
	"database/sql"

	"cpsu/internal/auth/models"
)

type AuthRepository interface {
	CreateUser(user models.Users) error
	GetByEmail(email string) (*models.Users, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(user models.Users) error {
	query := `
		INSERT INTO users (username, email, password_hash, is_active, email_verified)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(
		query, user.Username, user.Email,
		user.PasswordHash, user.IsActive,
		user.EmailVerified,
	)
	return err
}

func (r *authRepository) GetByEmail(email string) (*models.Users, error) {
	query := `
		SELECT user_id, username, email, password_hash, is_active, email_verified
		FROM users
		WHERE email = $1
	`

	var user models.Users
	err := r.db.QueryRow(query, email).Scan(
		&user.UserID, &user.Username, &user.Email,
		&user.PasswordHash, &user.IsActive,
		&user.EmailVerified,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
