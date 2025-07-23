package service

import (
	"cpsu/internal/models"
	"cpsu/internal/repository"
	"errors"
	"strings"
)

type NewsQueryParam struct {
	Search   string `json:"search"`
	Limit    int    `json:"limit"`
	NewsID   int    `json:"news_id"`
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
	NewsType string `json:"news_type"`
}

type CPSUService interface {
	GetAllNews(param NewsQueryParam) ([]models.News, error)
	GetNewsDetail(id int) (*models.News, error)
	CreateNews(title, content, newsType, detailURL string, images []string) (*models.News, error)
	UpdateNews(id int, title, content, newsType, detailURL string, images []string) (*models.News, error)
	DeleteNews(id int) error
}

type cpsuService struct {
	repo repository.CPSURepository
}

func NewCPSUService(repo repository.CPSURepository) CPSUService {
	return &cpsuService{repo: repo}
}

func (s *cpsuService) GetAllNews(param NewsQueryParam) ([]models.News, error) {
	param.Sort = "create_at"
	param.Order = "desc"
	return s.repo.GetAllNews(param)
}

func (s *cpsuService) GetNewsDetail(id int) (*models.News, error) {
	return s.repo.GetNewsDetail(id)
}

func (s *cpsuService) CreateNews(title, content, newsType, detailURL string, images []string) (*models.News, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		return nil, errors.New("title and content are required")
	}

	news := &models.News{
		Title:     title,
		Content:   content,
		NewsType:  newsType,
		DetailURL: detailURL,
	}

	created, err := s.repo.CreateNews(news)
	if err != nil {
		return nil, err
	}

	if len(images) > 0 {
		err = s.repo.AddNewsImages(created.NewsID, images)
		if err != nil {
			return nil, err
		}
	}

	return s.repo.GetNewsDetail(created.NewsID)
}

func (s *cpsuService) UpdateNews(id int, title, content, newsType, detailURL string, images []string) (*models.News, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		return nil, errors.New("title and content are required")
	}

	news := &models.News{
		Title:     title,
		Content:   content,
		NewsType:  newsType,
		DetailURL: detailURL,
	}

	updated, err := s.repo.UpdateNews(id, news)
	if err != nil {
		return nil, err
	}

	if len(images) > 0 {
		err = s.repo.ReplaceNewsImages(id, images)
		if err != nil {
			return nil, err
		}
		updated.Images = []models.NewsImage{}
		for _, img := range images {
			updated.Images = append(updated.Images, models.NewsImage{ImageURL: img})
		}
	}

	return updated, nil
}

func (s *cpsuService) DeleteNews(id int) error {
	return s.repo.DeleteNews(id)
}
