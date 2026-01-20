package service

import (
	"cpsu/internal/auth/models"
	"cpsu/internal/auth/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetAllUser(param models.UserQueryParam) ([]models.UserResponse, error) {
	return s.UserRepo.GetAllUser(param)
}
