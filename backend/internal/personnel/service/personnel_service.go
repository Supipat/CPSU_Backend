package service

import (
	"cpsu/internal/personnel/models"
	"cpsu/internal/personnel/repository"
)

type PersonnelService interface {
	GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error)
	GetPersonnelByID(id int) (*models.Personnels, error)
	CreatePersonnel(req models.PersonnelRequest) (*models.Personnels, error)
	UpdatePersonnel(id int, req models.PersonnelRequest) (*models.Personnels, error)
	DeletePersonnel(id int) error
}

type personnelService struct {
	repo repository.PersonnelRepository
}

func NewPersonnelService(repo repository.PersonnelRepository) PersonnelService {
	return &personnelService{
		repo: repo,
	}
}

func (s *personnelService) GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error) {
	return s.repo.GetAllPersonnels(param)
}

func (s *personnelService) GetPersonnelByID(id int) (*models.Personnels, error) {
	return s.repo.GetPersonnelByID(id)
}

func (s *personnelService) CreatePersonnel(req models.PersonnelRequest) (*models.Personnels, error) {
	return s.repo.CreatePersonnel(req)
}

func (s *personnelService) UpdatePersonnel(id int, req models.PersonnelRequest) (*models.Personnels, error) {
	return s.repo.UpdatePersonnel(id, req)
}

func (s *personnelService) DeletePersonnel(id int) error {
	return s.repo.DeletePersonnel(id)
}
