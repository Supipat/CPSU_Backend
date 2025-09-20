package service

import (
	"cpsu/internal/personnel/models"
	"cpsu/internal/personnel/repository"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type PersonnelService interface {
	GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error)
	GetPersonnelByID(id int) (*models.Personnels, error)
	CreatePersonnel(req models.PersonnelRequest, fileImage *multipart.FileHeader) (*models.Personnels, error)
	UpdatePersonnel(id int, req models.PersonnelRequest, fileImage *multipart.FileHeader) (*models.Personnels, error)
	DeletePersonnel(id int) error
}

type personnelService struct {
	repo   repository.PersonnelRepository
	upload *s3manager.Uploader
	bucket string
}

func NewPersonnelService(repo repository.PersonnelRepository, awsRegion, awsAccessKeyID, awsSecretAccessKey, bucket string) PersonnelService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))
	uploader := s3manager.NewUploader(sess)

	return &personnelService{
		repo:   repo,
		upload: uploader,
		bucket: bucket,
	}
}

func (s *personnelService) GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error) {
	return s.repo.GetAllPersonnels(param)
}

func (s *personnelService) GetPersonnelByID(id int) (*models.Personnels, error) {
	return s.repo.GetPersonnelByID(id)
}

func (s *personnelService) CreatePersonnel(req models.PersonnelRequest, fileImage *multipart.FileHeader) (*models.Personnels, error) {
	if fileImage != nil {
		url, err := s.UploadFile(fileImage)
		if err != nil {
			return nil, err
		}
		req.FileImage = url
	}
	return s.repo.CreatePersonnel(req)
}

func (s *personnelService) UpdatePersonnel(id int, req models.PersonnelRequest, fileImage *multipart.FileHeader) (*models.Personnels, error) {
	if fileImage != nil {
		url, err := s.UploadFile(fileImage)
		if err != nil {
			return nil, err
		}
		req.FileImage = url
	}
	return s.repo.UpdatePersonnel(id, req)
}

func (s *personnelService) DeletePersonnel(id int) error {
	return s.repo.DeletePersonnel(id)
}

func (s *personnelService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := "images/personnel/" + fileHeader.Filename

	input := &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   file,
	}

	result, err := s.upload.Upload(input)
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
