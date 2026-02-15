package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"cpsu/internal/admission/models"
	"cpsu/internal/admission/repository"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type AdmissionService interface {
	GetAllAdmission(param models.AdmissionQueryParam) ([]models.Admission, error)
	GetAdmissionByID(id int) (*models.Admission, error)
	CreateAdmission(req models.AdmissionRequest, fileImage *multipart.FileHeader) (*models.Admission, error)
	UpdateAdmission(id int, req models.AdmissionRequest, fileImage *multipart.FileHeader) (*models.Admission, error)
	DeleteAdmission(id int) error
}

type admissionService struct {
	repo        repository.AdmissionRepository
	minioClient *minio.Client
	bucket      string
	publicBase  string
}

func NewAdmissionService(
	repo repository.AdmissionRepository,
	endpoint string,
	accessKey string,
	secretKey string,
	bucket string,
	useSSL bool,
	publicBaseURL string,
) AdmissionService {

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	return &admissionService{
		repo:        repo,
		minioClient: client,
		bucket:      bucket,
		publicBase:  publicBaseURL,
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
		url, err := s.uploadFile(fileImage)
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
		url, err := s.uploadFile(fileImage)
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

func (s *admissionService) uploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	objectName := fmt.Sprintf(
		"admission/%s%s",
		uuid.New().String(),
		ext,
	)

	_, err = s.minioClient.PutObject(
		context.Background(),
		s.bucket,
		objectName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: fileHeader.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", err
	}

	// URL ถาวร (Public Bucket)
	imageURL := fmt.Sprintf(
		"%s/%s/%s",
		s.publicBase,
		s.bucket,
		objectName,
	)

	return imageURL, nil
}
