package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"

	authrepo "cpsu/internal/auth/repository"
	"cpsu/internal/document/models"
	documentrepo "cpsu/internal/document/repository"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type DocumentService interface {
	GetAllDocument(param models.DocumentQueryParam) ([]models.Document, error)
	GetDocumentByID(id int) (*models.Document, error)
	CreateDocument(typeID int, typeName string, title string, description *string, file *multipart.FileHeader, userID int, ip string, userAgent string) (*models.Document, error)
	UpdateDocument(id int, typeID int, typeName string, title string, description *string, file *multipart.FileHeader, userID int, ip string, userAgent string) (*models.Document, error)
	DeleteDocument(id int, userID int, ip string, userAgent string) error
}

type documentService struct {
	repo        documentrepo.DocumentRepository
	auditRepo   *authrepo.AuditRepository
	minioClient *minio.Client
	bucket      string
	publicBase  string
}

func NewDocumentService(repo documentrepo.DocumentRepository, auditRepo *authrepo.AuditRepository, endpoint string, accessKey string, secretKey string, bucket string, useSSL bool, publicBaseURL string) DocumentService {

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	return &documentService{
		repo:        repo,
		auditRepo:   auditRepo,
		minioClient: client,
		bucket:      bucket,
		publicBase:  publicBaseURL,
	}
}

func (s *documentService) GetAllDocument(param models.DocumentQueryParam) ([]models.Document, error) {

	if param.Sort == "" {
		param.Sort = "document_id"
	}

	if param.Order == "" {
		param.Order = "desc"
	}

	documents, err := s.repo.GetAllDocument(param)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (s *documentService) GetDocumentByID(id int) (*models.Document, error) {
	return s.repo.GetDocumentByID(id)
}

func (s *documentService) CreateDocument(typeID int, typeName string, title string, description *string, file *multipart.FileHeader, userID int, ip string, userAgent string) (*models.Document, error) {

	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title is required")
	}

	if file == nil {
		return nil, errors.New("file is required")
	}

	fileURL, err := s.UploadFile(file)
	if err != nil {
		return nil, err
	}

	req := models.DocumentRequest{TypeID: typeID, TypeName: typeName, Title: title, Description: description, File: fileURL}

	created, err := s.repo.CreateDocument(req)
	if err != nil {
		return nil, err
	}

	_ = s.auditRepo.LogAudit(
		userID,
		"create",
		"document",
		strconv.Itoa(created.DocumentID),
		map[string]interface{}{
			"title": created.Title,
		},
		ip,
		userAgent,
	)

	return s.repo.GetDocumentByID(created.DocumentID)
}

func (s *documentService) UpdateDocument(id int, typeID int, typeName string, title string, description *string, file *multipart.FileHeader, userID int, ip string, userAgent string) (*models.Document, error) {

	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title is required")
	}

	oldDocument, err := s.repo.GetDocumentByID(id)
	if err != nil {
		return nil, err
	}

	fileURL := oldDocument.File

	if file != nil {
		uploadedURL, err := s.UploadFile(file)
		if err != nil {
			return nil, err
		}

		fileURL = uploadedURL
	}

	req := models.DocumentRequest{TypeID: typeID, TypeName: typeName, Title: title, Description: description, File: fileURL}

	updated, err := s.repo.UpdateDocument(id, req)
	if err != nil {
		return nil, err
	}

	_ = s.auditRepo.LogAudit(
		userID,
		"update",
		"document",
		strconv.Itoa(updated.DocumentID),
		map[string]interface{}{
			"title": updated.Title,
		},
		ip,
		userAgent,
	)

	return s.repo.GetDocumentByID(updated.DocumentID)
}

func (s *documentService) DeleteDocument(id int, userID int, ip string, userAgent string) error {

	err := s.repo.DeleteDocument(id)
	if err != nil {
		return err
	}

	_ = s.auditRepo.LogAudit(
		userID,
		"delete",
		"document",
		strconv.Itoa(id),
		map[string]interface{}{},
		ip,
		userAgent,
	)

	return nil
}

func (s *documentService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)

	objectName := fmt.Sprintf(
		"document/%s%s",
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

	fileURL := fmt.Sprintf(
		"%s/%s/%s",
		s.publicBase,
		s.bucket,
		objectName,
	)

	return fileURL, nil
}