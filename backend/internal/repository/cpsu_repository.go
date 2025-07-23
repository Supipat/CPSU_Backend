package repository

import (
	"cpsu/internal/models"
	"database/sql"

	"github.com/lib/pq"
)

type CPSURepository interface {
	GetAllNews(param interface{}) ([]models.News, error)
	GetNewsDetail(id int) (*models.News, error)
	CreateNews(news *models.News) (*models.News, error)
	UpdateNews(id int, news *models.News) (*models.News, error)
	DeleteNews(id int) error
	AddNewsImages(newsID int, images []string) error
	ReplaceNewsImages(newsID int, images []string) error
}

type cpsuRepository struct {
	db *sql.DB
}

func NewCPSURepository(db *sql.DB) CPSURepository {
	return &cpsuRepository{db: db}
}

func (r *cpsuRepository) GetAllNews(param interface{}) ([]models.News, error) {
	query := `
		SELECT n.news_id, n.title, n.content, n.news_type,
		       n.detail_url, n.create_at, n.updated_at,
		       COALESCE(array_agg(ni.image_id) FILTER (WHERE ni.image_id IS NOT NULL), '{}'),
		       COALESCE(array_agg(ni.image_url) FILTER (WHERE ni.image_url IS NOT NULL), '{}')
		FROM news n
		LEFT JOIN news_image ni ON n.news_id = ni.news_id
		GROUP BY n.news_id
		ORDER BY n.create_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allNews []models.News
	for rows.Next() {
		var news models.News
		var imageIDs []sql.NullInt64
		var imageURLs []sql.NullString

		err := rows.Scan(
			&news.NewsID, &news.Title, &news.Content, &news.NewsType,
			&news.DetailURL, &news.CreatedAt, &news.UpdatedAt,
			pq.Array(&imageIDs), pq.Array(&imageURLs),
		)
		if err != nil {
			return nil, err
		}

		for i := range imageIDs {
			if imageIDs[i].Valid && imageURLs[i].Valid {
				news.Images = append(news.Images, models.NewsImage{
					ImageID:  int(imageIDs[i].Int64),
					NewsID:   news.NewsID,
					ImageURL: imageURLs[i].String,
				})
			}
		}
		allNews = append(allNews, news)
	}
	return allNews, nil
}

func (r *cpsuRepository) GetNewsDetail(id int) (*models.News, error) {
	query := `
		SELECT n.news_id, n.title, n.content, n.news_type,
		       n.detail_url, n.create_at, n.updated_at,
		       COALESCE(array_agg(ni.image_id) FILTER (WHERE ni.image_id IS NOT NULL), '{}'),
		       COALESCE(array_agg(ni.image_url) FILTER (WHERE ni.image_url IS NOT NULL), '{}')
		FROM news n
		LEFT JOIN news_image ni ON n.news_id = ni.news_id
		WHERE n.news_id = $1
		GROUP BY n.news_id
	`
	row := r.db.QueryRow(query, id)

	var news models.News
	var imageIDs []sql.NullInt64
	var imageURLs []sql.NullString

	err := row.Scan(
		&news.NewsID, &news.Title, &news.Content, &news.NewsType,
		&news.DetailURL, &news.CreatedAt, &news.UpdatedAt,
		pq.Array(&imageIDs), pq.Array(&imageURLs),
	)
	if err != nil {
		return nil, err
	}

	for i := range imageIDs {
		if imageIDs[i].Valid && imageURLs[i].Valid {
			news.Images = append(news.Images, models.NewsImage{
				ImageID:  int(imageIDs[i].Int64),
				NewsID:   news.NewsID,
				ImageURL: imageURLs[i].String,
			})
		}
	}
	return &news, nil
}

func (r *cpsuRepository) CreateNews(news *models.News) (*models.News, error) {
	query := `
		INSERT INTO news (title, content, news_type, detail_url)
		VALUES ($1, $2, $3, $4)
		RETURNING news_id, create_at, updated_at
	`
	err := r.db.QueryRow(
		query,
		news.Title, news.Content, news.NewsType, news.DetailURL,
	).Scan(&news.NewsID, &news.CreatedAt, &news.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *cpsuRepository) UpdateNews(id int, news *models.News) (*models.News, error) {
	query := `
		UPDATE news
		SET title = $1, content = $2, news_type = $3, detail_url = $4, updated_at = NOW()
		WHERE news_id = $5
		RETURNING news_id, create_at, updated_at
	`
	err := r.db.QueryRow(
		query,
		news.Title, news.Content, news.NewsType, news.DetailURL, id,
	).Scan(&news.NewsID, &news.CreatedAt, &news.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *cpsuRepository) DeleteNews(id int) error {
	result, err := r.db.Exec("DELETE FROM news WHERE news_id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *cpsuRepository) AddNewsImages(newsID int, images []string) error {
	for _, img := range images {
		_, err := r.db.Exec("INSERT INTO news_image (news_id, image_url) VALUES ($1, $2)", newsID, img)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *cpsuRepository) ReplaceNewsImages(newsID int, images []string) error {
	_, err := r.db.Exec("DELETE FROM news_image WHERE news_id = $1", newsID)
	if err != nil {
		return err
	}
	return r.AddNewsImages(newsID, images)
}
