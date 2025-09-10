package service

import (
	"cpsu/internal/course/models"
	"cpsu/internal/course/repository"
)

type CourseService interface {
	GetAllCourses(param models.CoursesQueryParam) ([]models.Courses, error)
	GetCourseByID(id string) (*models.Courses, error)
	CreateCourse(req models.CoursesRequest) (*models.Courses, error)
	UpdateCourse(id string, req models.CoursesRequest) (*models.Courses, error)
	DeleteCourse(id string) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{
		repo: repo,
	}
}

func (s *courseService) GetAllCourses(param models.CoursesQueryParam) ([]models.Courses, error) {
	return s.repo.GetAllCourses(param)
}

func (s *courseService) GetCourseByID(id string) (*models.Courses, error) {
	return s.repo.GetCourseByID(id)
}

func (s *courseService) CreateCourse(course models.CoursesRequest) (*models.Courses, error) {
	return s.repo.CreateCourse(course)
}

func (s *courseService) UpdateCourse(id string, course models.CoursesRequest) (*models.Courses, error) {
	return s.repo.UpdateCourse(id, course)
}

func (s *courseService) DeleteCourse(id string) error {
	return s.repo.DeleteCourse(id)
}
