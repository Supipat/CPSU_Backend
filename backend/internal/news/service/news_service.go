package service

import (
	"errors"
	"strings"

	"cpsu/internal/news/models"
	"cpsu/internal/news/repository"
)

type CPSUService interface {
	GetAllNews(param models.NewsQueryParam) ([]models.News, error)
	GetNewsByID(id int) (*models.News, error)
	CreateNews(title, content string, typeID int, typeName, detailURL string, images []string) (*models.News, error)
	UpdateNews(id int, title, content string, type_id int, typeName, detailURL string, images []string) (*models.News, error)
	DeleteNews(id int) error
}

type cpsuService struct {
	repo repository.CPSURepository
}

func NewCPSUService(repo repository.CPSURepository) CPSUService {
	return &cpsuService{repo: repo}
}

func (s *cpsuService) GetAllNews(param models.NewsQueryParam) ([]models.News, error) {
	if param.Sort == "" {
		param.Sort = "created_at"
	}
	if param.Order == "" {
		param.Order = "desc"
	}
	newsList, err := s.repo.GetAllNews(param)
	if err != nil {
		return nil, err
	}
	if len(newsList) == 0 && param.TypeName != "" {
		return nil, errors.New("news type not found")
	}

	return newsList, nil
}

func (s *cpsuService) GetNewsByID(id int) (*models.News, error) {
	return s.repo.GetNewsByID(id)
}

func (s *cpsuService) CreateNews(title, content string, typeID int, typeName, detailURL string, images []string) (*models.News, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		return nil, errors.New("title and content are required")
	}

	newsreq := &models.NewsRequest{
		Title:     title,
		Content:   content,
		TypeID:    typeID,
		DetailURL: detailURL,
	}

	created, err := s.repo.CreateNews(newsreq)
	if err != nil {
		return nil, err
	}

	if len(images) > 0 {
		err = s.repo.AddNewsImages(created.NewsID, images)
		if err != nil {
			return nil, err
		}
	}

	return s.repo.GetNewsByID(created.NewsID)
}

func (s *cpsuService) UpdateNews(id int, title, content string, typeID int, typeName, detailURL string, images []string) (*models.News, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		return nil, errors.New("title and content are required")
	}

	newsreq := &models.NewsRequest{
		Title:     title,
		Content:   content,
		TypeID:    typeID,
		DetailURL: detailURL,
	}

	_, err := s.repo.UpdateNews(id, newsreq)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.UpdateNewsImages(id, images)
	if err != nil {
		return nil, err
	}

	return s.repo.GetNewsByID(id)
}

func (s *cpsuService) DeleteNews(id int) error {
	return s.repo.DeleteNews(id)
}
