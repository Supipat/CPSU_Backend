package service

import (
	"errors"
	"mime/multipart"
	"strings"

	"cpsu/internal/course_structure/models"
	"cpsu/internal/course_structure/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type CourseStructureService interface {
	GetAllCourseStructure(param models.CourseStructureQueryParam) ([]models.CourseStructure, error)
	GetCourseStructureByID(id int) (*models.CourseStructure, error)
	CreateCourseStructure(courseID int, file *multipart.FileHeader) (*models.CourseStructure, error)
	DeleteCourseStructure(id int) error
}

type courseStructureService struct {
	repo   repository.CourseStructureRepository
	upload *s3manager.Uploader
	bucket string
}

func NewCourseStructureService(repo repository.CourseStructureRepository, awsRegion, awsAccessKeyID, awsSecretAccessKey, bucket string) CourseStructureService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))
	upload := s3manager.NewUploader(sess)

	return &courseStructureService{
		repo:   repo,
		upload: upload,
		bucket: bucket,
	}
}

func (s *courseStructureService) GetAllCourseStructure(param models.CourseStructureQueryParam) ([]models.CourseStructure, error) {
	return s.repo.GetAllCourseStructure(param)
}

func (s *courseStructureService) GetCourseStructureByID(id int) (*models.CourseStructure, error) {
	return s.repo.GetCourseStructureByID(id)
}

func (s *courseStructureService) CreateCourseStructure(courseID int, file *multipart.FileHeader) (*models.CourseStructure, error) {
	if courseID == 0 {
		return nil, errors.New("course_id is required")
	}
	if file == nil {
		return nil, errors.New("course structure file is required")
	}

	url, err := s.UploadFile(file)
	if err != nil {
		return nil, err
	}

	req := &models.CourseStructureRequest{
		CourseID:           courseID,
		CourseStructureURL: url,
	}

	return s.repo.CreateCourseStructure(req)
}

func (s *courseStructureService) DeleteCourseStructure(id int) error {
	return s.repo.DeleteCourseStructure(id)
}

func (s *courseStructureService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
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
