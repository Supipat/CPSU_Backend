package service

import (
	"cpsu/internal/personnel/models"
	"cpsu/internal/personnel/repository"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type PersonnelService interface {
	GetAllPersonnels(param models.PersonnelQueryParam) ([]models.Personnels, error)
	GetPersonnelByID(id int) (*models.Personnels, error)
	GetResearchByScopusID(scopusID string) ([]models.Research, error)
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
	personnels, err := s.repo.GetAllPersonnels(param)
	if err != nil {
		return nil, err
	}

	for i := range personnels {
		if personnels[i].ScopusID != nil && *personnels[i].ScopusID != "" {
			research, err := s.GetResearchByScopusID(*personnels[i].ScopusID)
			if err == nil {
				personnels[i].Researches = research
			}
		}
	}

	return personnels, nil
}

func (s *personnelService) GetPersonnelByID(id int) (*models.Personnels, error) {
	return s.repo.GetPersonnelByID(id)
}

func (s *personnelService) GetResearchByScopusID(scopusID string) ([]models.Research, error) {
	apiKey := os.Getenv("SCOPUS_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing scopus_api_key")
	}

	url := fmt.Sprintf(
		"https://api.elsevier.com/content/search/scopus?query=AU-ID(%s)&apiKey=%s",
		scopusID, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to request scopus api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("scopus api returned status: %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode scopus response: %w", err)
	}

	results, _ := data["search-results"].(map[string]interface{})
	entries, _ := results["entry"].([]interface{})

	var list []models.Research
	for _, e := range entries {
		item := e.(map[string]interface{})
		list = append(list, models.Research{
			Title:   fmt.Sprint(item["dc:title"]),
			Journal: fmt.Sprint(item["prism:publicationName"]),
			Year:    fmt.Sprint(item["prism:coverDate"]),
			DOI:     fmt.Sprint(item["prism:doi"]),
			Cited:   fmt.Sprint(item["citedby-count"]),
		})
	}

	return list, nil
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
