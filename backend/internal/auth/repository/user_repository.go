package repository

import (
	"database/sql"
	"errors"

	"cpsu/internal/auth/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Register(username, email, passwordHash string) (int, error) {
	query := `
		INSERT INTO users (username, email, password_hash, is_active)
		VALUES ($1, $2, $3, true)
		RETURNING user_id
	`

	var userID int
	err := r.DB.QueryRow(
		query, username, email, passwordHash,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `
		SELECT user_id, username, email, password_hash, is_active, created_at, last_login
		FROM users
		WHERE email = $1
	`

	var user models.User
	var lastLogin sql.NullTime

	err := r.DB.QueryRow(query, email).Scan(
		&user.UserID, &user.Username, &user.Email, &user.PasswordHash,
		&user.IsActive, &user.CreatedAt, &lastLogin,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

func (r *UserRepository) FindByID(userID int) (*models.User, error) {
	query := `
		SELECT
			user_id, username, email, password_hash, is_active, created_at, last_login
		FROM users
		WHERE user_id = $1
	`

	var user models.User
	var lastLogin sql.NullTime

	err := r.DB.QueryRow(query, userID).Scan(
		&user.UserID, &user.Username, &user.Email, &user.PasswordHash,
		&user.IsActive, &user.CreatedAt, &lastLogin,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

func (r *UserRepository) UpdateLastLogin(userID int) error {
	query := `
		UPDATE users
		SET last_login = NOW(),
		    updated_at = NOW()
		WHERE user_id = $1
	`
	_, err := r.DB.Exec(query, userID)
	return err
}
