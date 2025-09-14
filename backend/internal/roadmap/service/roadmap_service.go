package service

import (
	"errors"
	"mime/multipart"
	"strings"

	"cpsu/internal/roadmap/models"
	"cpsu/internal/roadmap/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type RoadmapService interface {
	GetAllRoadmap(param models.RoadmapQueryParam) ([]models.Roadmap, error)
	GetRoadmapByID(id int) (*models.Roadmap, error)
	CreateRoadmap(courseID string, file *multipart.FileHeader) (*models.Roadmap, error)
	DeleteRoadmap(id int) error
}

type roadmapService struct {
	repo   repository.RoadmapRepository
	upload *s3manager.Uploader
	bucket string
}

func NewRoadmapService(repo repository.RoadmapRepository, awsRegion, awsAccessKeyID, awsSecretAccessKey, bucket string) RoadmapService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))
	upload := s3manager.NewUploader(sess)

	return &roadmapService{
		repo:   repo,
		upload: upload,
		bucket: bucket,
	}
}

func (s *roadmapService) GetAllRoadmap(param models.RoadmapQueryParam) ([]models.Roadmap, error) {
	return s.repo.GetAllRoadmap(param)
}

func (s *roadmapService) GetRoadmapByID(id int) (*models.Roadmap, error) {
	return s.repo.GetRoadmapByID(id)
}

func (s *roadmapService) CreateRoadmap(courseID string, file *multipart.FileHeader) (*models.Roadmap, error) {
	if courseID == "" {
		return nil, errors.New("course_id is required")
	}
	if file == nil {
		return nil, errors.New("roadmap image is required")
	}

	url, err := s.UploadFile(file)
	if err != nil {
		return nil, err
	}

	req := &models.RoadmapRequest{
		CourseID:   courseID,
		RoadmapURL: url,
	}

	return s.repo.CreateRoadmap(req)
}

func (s *roadmapService) DeleteRoadmap(id int) error {
	return s.repo.DeleteRoadmap(id)
}

func (s *roadmapService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileName := strings.ReplaceAll(fileHeader.Filename, " ", "_")
	key := "images/course/" + fileName

	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   file,
	}

	result, err := s.upload.Upload(uploadInput)
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
