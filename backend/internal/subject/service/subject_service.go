package service

import (
	"cpsu/internal/subject/models"
	"cpsu/internal/subject/repository"
)

type SubjectService interface {
	GetAllSubjects(param models.SubjectsQueryParam) ([]models.Subjects, error)
	GetSubjectByID(id int) (*models.Subjects, error)
	CreateSubject(req models.SubjectsRequest) (*models.Subjects, error)
	UpdateSubject(id int, req models.SubjectsRequest) (*models.Subjects, error)
	DeleteSubject(id int) error
}

type subjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) SubjectService {
	return &subjectService{
		repo: repo,
	}
}

func (s *subjectService) GetAllSubjects(param models.SubjectsQueryParam) ([]models.Subjects, error) {
	return s.repo.GetAllSubjects(param)
}

func (s *subjectService) GetSubjectByID(id int) (*models.Subjects, error) {
	return s.repo.GetSubjectByID(id)
}

func (s *subjectService) CreateSubject(subject models.SubjectsRequest) (*models.Subjects, error) {
	return s.repo.CreateSubject(subject)
}

func (s *subjectService) UpdateSubject(id int, subject models.SubjectsRequest) (*models.Subjects, error) {
	return s.repo.UpdateSubject(id, subject)
}

func (s *subjectService) DeleteSubject(id int) error {
	return s.repo.DeleteSubject(id)
}
