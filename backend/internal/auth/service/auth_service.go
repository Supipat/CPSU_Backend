package service

import (
	"errors"

	"cpsu/internal/auth/models"
	"cpsu/internal/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repository.AuthRepository
}

const bcryptCost = 12

func NewAuthService(userRepo repository.AuthRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(req models.RegisterRequest) error {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return errors.New("missing required fields")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return err
	}

	user := models.Users{
		Username:      req.Username,
		Email:         req.Email,
		PasswordHash:  string(hash),
		EmailVerified: false,
		IsActive:      true,
	}

	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(req models.LoginRequest) (*models.Users, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password required")
	}

	user, err := s.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is disabled")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
