package service

import (
	"cpsu/internal/admission/models"
	"cpsu/internal/admission/repository"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AdmissionService interface {
	GetAllAdmission(param models.AdmissionQueryParam) ([]models.Admission, error)
	GetAdmissionByID(id int) (*models.Admission, error)
	CreateAdmission(req models.AdmissionRequest, fileImage *multipart.FileHeader) (*models.Admission, error)
	UpdateAdmission(id int, req models.AdmissionRequest, fileImage *multipart.FileHeader) (*models.Admission, error)
	DeleteAdmission(id int) error
}

type admissionService struct {
	repo   repository.AdmissionRepository
	upload *s3manager.Uploader
	bucket string
}

func NewAdmissionService(
	repo repository.AdmissionRepository,
	awsRegion,
	awsAccessKeyID,
	awsSecretAccessKey,
	bucket string,
) AdmissionService {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))

	uploader := s3manager.NewUploader(sess)

	return &admissionService{
		repo:   repo,
		upload: uploader,
		bucket: bucket,
	}
}

func (s *admissionService) GetAllAdmission(param models.AdmissionQueryParam) ([]models.Admission, error) {
	return s.repo.GetAllAdmission(param)
}

func (s *admissionService) GetAdmissionByID(id int) (*models.Admission, error) {
	return s.repo.GetAdmissionByID(id)
}

func (s *admissionService) CreateAdmission(
	req models.AdmissionRequest,
	fileImage *multipart.FileHeader,
) (*models.Admission, error) {

	if fileImage != nil {
		url, err := s.UploadFile(fileImage)
		if err != nil {
			return nil, err
		}
		req.FileImage = url
	}

	return s.repo.CreateAdmission(req)
}

func (s *admissionService) UpdateAdmission(
	id int,
	req models.AdmissionRequest,
	fileImage *multipart.FileHeader,
) (*models.Admission, error) {

	if fileImage != nil {
		url, err := s.UploadFile(fileImage)
		if err != nil {
			return nil, err
		}
		req.FileImage = url
	}

	return s.repo.UpdateAdmission(id, req)
}

func (s *admissionService) DeleteAdmission(id int) error {
	return s.repo.DeleteAdmission(id)
}

func (s *admissionService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := "images/admission/" + fileHeader.Filename

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
