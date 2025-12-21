package service

import (
	"errors"
	"strings"

	"cpsu/internal/course_structure/models"
	"cpsu/internal/course_structure/repository"
)

type CourseStructureService interface {
	GetAllCourseStructure(param models.CourseStructureQueryParam) ([]models.CourseStructure, error)
	GetCourseStructureByID(id int) (*models.CourseStructure, error)
	CreateCourseStructure(req models.CourseStructureRequest) (*models.CourseStructure, error)
	UpdateCourseStructure(id int, req models.CourseStructureRequest) (*models.CourseStructure, error)
	DeleteCourseStructure(id int) error
}

type courseStructureService struct {
	repo repository.CourseStructureRepository
}

func NewCourseStructureService(
	repo repository.CourseStructureRepository,
) CourseStructureService {
	return &courseStructureService{
		repo: repo,
	}
}

func (s *courseStructureService) GetAllCourseStructure(param models.CourseStructureQueryParam) ([]models.CourseStructure, error) {
	return s.repo.GetAllCourseStructure(param)
}

func (s *courseStructureService) GetCourseStructureByID(id int) (*models.CourseStructure, error) {
	if id <= 0 {
		return nil, errors.New("invalid course_structure_id")
	}

	return s.repo.GetCourseStructureByID(id)
}

func (s *courseStructureService) CreateCourseStructure(req models.CourseStructureRequest) (*models.CourseStructure, error) {
	if strings.TrimSpace(req.CourseID) == "" {
		return nil, errors.New("course_id is required")
	}

	if strings.TrimSpace(req.Detail) == "" {
		return nil, errors.New("detail is required")
	}

	return s.repo.CreateCourseStructure(&req)
}

func (s *courseStructureService) UpdateCourseStructure(id int, req models.CourseStructureRequest) (*models.CourseStructure, error) {
	return s.repo.UpdateCourseStructure(id, req)
}

func (s *courseStructureService) DeleteCourseStructure(id int) error {
	if id <= 0 {
		return errors.New("invalid course_structure_id")
	}

	return s.repo.DeleteCourseStructure(id)
}
