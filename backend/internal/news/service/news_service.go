package service

import (
	"errors"
	"mime/multipart"
	"strings"

	"cpsu/internal/news/models"
	"cpsu/internal/news/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type NewsService interface {
	GetAllNews(param models.NewsQueryParam) ([]models.News, error)
	GetNewsByID(id int) (*models.News, error)
	CreateNews(title, content string, typeID int, typeName, detailURL string, coverImage *multipart.FileHeader, images []*multipart.FileHeader) (*models.News, error)
	UpdateNews(id int, title, content string, type_id int, typeName, detailURL string, coverImage *multipart.FileHeader, images []*multipart.FileHeader) (*models.News, error)
	DeleteNews(id int) error
}

type newsService struct {
	repo   repository.NewsRepository
	upload *s3manager.Uploader
	bucket string
}

func NewNewsService(repo repository.NewsRepository, awsRegion, awsAccessKeyID, awsSecretAccessKey, bucket string) NewsService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))
	upload := s3manager.NewUploader(sess)

	return &newsService{
		repo:   repo,
		upload: upload,
		bucket: bucket,
	}
}

func (s *newsService) GetAllNews(param models.NewsQueryParam) ([]models.News, error) {
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
	if len(newsList) == 0 && param.TypeID > 0 {
		return nil, errors.New("news type not found")
	}

	return newsList, nil
}

func (s *newsService) GetNewsByID(id int) (*models.News, error) {
	return s.repo.GetNewsByID(id)
}

func (s *newsService) CreateNews(title, content string, typeID int, typeName, detailURL string, coverImage *multipart.FileHeader, images []*multipart.FileHeader) (*models.News, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		return nil, errors.New("title and content are required")
	}

	var uploadedFlies []string
	for _, fileHeader := range images {
		url, err := s.UploadImages(fileHeader)
		if err != nil {
			return nil, err
		}
		uploadedFlies = append(uploadedFlies, url)
	}

	var coverURL string
	if coverImage != nil {
		url, err := s.UploadImages(coverImage)
		if err != nil {
			return nil, err
		}
		coverURL = url
	}

	newsReq := &models.NewsRequest{
		Title:      title,
		Content:    content,
		TypeID:     typeID,
		DetailURL:  detailURL,
		CoverImage: coverURL,
	}

	for _, url := range uploadedFlies {
		newsReq.Images = append(newsReq.Images, models.NewsImages{FileImage: url})
	}

	created, err := s.repo.CreateNews(newsReq)
	if err != nil {
		return nil, err
	}

	return s.repo.GetNewsByID(created.NewsID)
}

func (s *newsService) UpdateNews(id int, title, content string, typeID int, typeName, detailURL string, coverImage *multipart.FileHeader, images []*multipart.FileHeader) (*models.News, error) {
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		return nil, errors.New("title and content are required")
	}

	var uploadedFlies []string
	for _, fileHeader := range images {
		url, err := s.UploadImages(fileHeader)
		if err != nil {
			return nil, err
		}
		uploadedFlies = append(uploadedFlies, url)
	}

	var coverURL string
	if coverImage != nil {
		url, err := s.UploadImages(coverImage)
		if err != nil {
			return nil, err
		}
		coverURL = url
	}

	newsReq := &models.NewsRequest{
		Title:      title,
		Content:    content,
		TypeID:     typeID,
		DetailURL:  detailURL,
		CoverImage: coverURL,
	}

	for _, url := range uploadedFlies {
		newsReq.Images = append(newsReq.Images, models.NewsImages{FileImage: url})
	}

	_, err := s.repo.UpdateNews(id, newsReq)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.UpdateNewsImages(id, repository.ImagesAsStrings(newsReq.Images))
	if err != nil {
		return nil, err
	}

	return s.repo.GetNewsByID(id)
}

func (s *newsService) DeleteNews(id int) error {
	return s.repo.DeleteNews(id)
}

func (s *newsService) UploadImages(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := "images/news/" + fileHeader.Filename

	newsImage := &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   file,
	}

	result, err := s.upload.Upload(newsImage)
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
