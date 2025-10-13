package service

import (
	"cpsu/internal/personnel/models"
	"cpsu/internal/personnel/repository"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

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
	GetResearchByPersonnelID(personnelID int) ([]models.Research, error)
	SyncResearch(personnelID int) ([]models.Research, error)
	GetResearchFromScopus(scopusID string) ([]models.Research, error)
	SyncAllFromScopus() (int, error)
	GetAllResearch(param models.ResearchQueryParam) ([]models.Research, error)
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

func (s *personnelService) GetResearchByPersonnelID(personnelID int) ([]models.Research, error) {
	return s.repo.GetResearchByPersonnelID(personnelID)
}

func (s *personnelService) SyncResearch(personnelID int) ([]models.Research, error) {
	scopusIDptr, err := s.repo.GetScopusIDByPersonnelID(personnelID)
	if err != nil {
		return nil, err
	}
	if scopusIDptr == nil || *scopusIDptr == "" {
		return nil, fmt.Errorf("personnel %d has no scopus_id", personnelID)
	}

	rs, err := s.GetResearchFromScopus(*scopusIDptr)
	if err != nil {
		return nil, err
	}
	for i := range rs {
		rs[i].PersonnelID = personnelID
		if rs[i].CreatedAt.IsZero() {
			rs[i].CreatedAt = time.Now()
		}
	}
	if err := s.repo.SaveResearch(personnelID, rs); err != nil {
		return nil, err
	}
	return s.repo.GetResearchByPersonnelID(personnelID)
}

func (s *personnelService) GetResearchFromScopus(scopusID string) ([]models.Research, error) {
	apiKey := os.Getenv("SCOPUS_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing SCOPUS_API_KEY")
	}
	url := fmt.Sprintf("https://api.elsevier.com/content/search/scopus?query=AU-ID(%s)&apiKey=%s", scopusID, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("scopus status %d", resp.StatusCode)
	}
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	sr, ok := data["search-results"].(map[string]interface{})
	if !ok {
		return nil, nil
	}
	entriesRaw := sr["entry"]
	entries, ok := entriesRaw.([]interface{})
	if !ok || len(entries) == 0 {
		return []models.Research{}, nil
	}

	toPtr := func(v interface{}) *string {
		if v == nil {
			return nil
		}
		s := fmt.Sprint(v)
		if s == "" {
			return nil
		}
		return &s
	}

	var research []models.Research
	for _, e := range entries {
		item, ok := e.(map[string]interface{})
		if !ok {
			continue
		}
		year := 0
		if ds, ok := item["prism:coverDate"].(string); ok && len(ds) >= 4 {
			if y, err := strconv.Atoi(ds[:4]); err == nil {
				year = y
			}
		}
		cited := 0
		switch v := item["citedby-count"].(type) {
		case string:
			cited, _ = strconv.Atoi(v)
		case float64:
			cited = int(v)
		}
		research = append(research, models.Research{
			Title:     fmt.Sprint(item["dc:title"]),
			Journal:   fmt.Sprint(item["prism:publicationName"]),
			Year:      year,
			Volume:    toPtr(item["prism:volume"]),
			Issue:     toPtr(item["prism:issueIdentifier"]),
			Pages:     toPtr(item["prism:pageRange"]),
			DOI:       toPtr(item["prism:doi"]),
			Cited:     cited,
			CreatedAt: time.Now(),
		})
	}
	return research, nil
}

func (s *personnelService) SyncAllFromScopus() (int, error) {
	personnels, err := s.repo.GetAllPersonnels(models.PersonnelQueryParam{})
	if err != nil {
		return 0, err
	}

	processed := 0

	fmt.Printf("SyncAllFromScopus: total personnels = %d\n", len(personnels))

	for _, p := range personnels {
		if p.ScopusID == nil || *p.ScopusID == "" {
			continue
		}
		rs, err := s.GetResearchFromScopus(*p.ScopusID)
		if err != nil {
			continue
		}
		if len(rs) == 0 {
			continue
		}

		sort.Slice(rs, func(i, j int) bool {
			return rs[i].Year > rs[j].Year
		})
		if len(rs) > 5 {
			rs = rs[:5]
		}

		for i := range rs {
			rs[i].PersonnelID = p.PersonnelID
			if rs[i].CreatedAt.IsZero() {
				rs[i].CreatedAt = time.Now()
			}
		}

		if err := s.repo.SaveResearch(p.PersonnelID, rs); err == nil {
			processed++
		}
	}
	return processed, nil
}

func (s *personnelService) GetAllResearch(param models.ResearchQueryParam) ([]models.Research, error) {
	return s.repo.GetAllResearch(param)
}
