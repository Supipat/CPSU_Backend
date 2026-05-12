package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"cpsu/internal/document/models"
)

type DocumentRepository interface {
	GetAllDocument(param models.DocumentQueryParam) ([]models.Document, error)
	GetDocumentByID(id int) (*models.Document, error)
	CreateDocument(req models.DocumentRequest) (*models.Document, error)
	UpdateDocument(id int, req models.DocumentRequest) (*models.Document, error)
	DeleteDocument(id int) error
}

type documentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) DocumentRepository {
	return &documentRepository{db: db}
}

func (r *documentRepository) GetAllDocument(param models.DocumentQueryParam) ([]models.Document, error) {
	query := `
		SELECT 
			d.document_id, d.type_id, dt.type_name, d.title, d.description, d.file
		FROM document d
		LEFT JOIN document_types dt ON d.type_id = dt.type_id
	`

	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if param.TypeID > 0 {
		conditions = append(conditions, "d.type_id = $"+strconv.Itoa(argIndex))
		args = append(args, param.TypeID)
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	sort := "d.document_id"
	if param.Sort != "" {
		sort = "d." + param.Sort
	}

	order := "ASC"
	if strings.ToUpper(param.Order) == "DESC" {
		order = "DESC"
	}

	query += " ORDER BY " + sort + " " + order

	if param.Limit > 0 {
		query += " LIMIT " + strconv.Itoa(param.Limit)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []models.Document

	for rows.Next() {
		var document models.Document

		err := rows.Scan(
			&document.DocumentID, &document.TypeID, &document.TypeName,
			&document.Title, &document.Description, &document.File,
		)
		if err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	return documents, nil
}

func (r *documentRepository) GetDocumentByID(id int) (*models.Document, error) {
	query := `
		SELECT 
			d.document_id, d.type_id, dt.type_name, d.title, d.description, d.file
		FROM document d
		LEFT JOIN document_types dt ON d.type_id = dt.type_id
		WHERE d.document_id = $1
	`

	row := r.db.QueryRow(query, id)

	var document models.Document

	err := row.Scan(
		&document.DocumentID, &document.TypeID, &document.TypeName,
		&document.Title, &document.Description, &document.File,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &document, nil
}

func (r *documentRepository) CreateDocument(req models.DocumentRequest) (*models.Document, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var typeID int

	if req.TypeID != 0 {
		typeID = req.TypeID
	} else if req.TypeName != "" {

		err = tx.QueryRow(`
			SELECT type_id FROM document_types WHERE type_name = $1
		`, req.TypeName).Scan(&typeID)

		if err == sql.ErrNoRows {

			err = tx.QueryRow(`
				INSERT INTO document_types (type_name) VALUES ($1) RETURNING type_id
			`, req.TypeName).Scan(&typeID)

			if err != nil {
				return nil, err
			}

		} else if err != nil {
			return nil, err
		}
	}

	var documentID int

	err = tx.QueryRow(`
		INSERT INTO document (type_id, title, description, file)
		VALUES ($1, $2, $3, $4)
		RETURNING document_id
	`,
		typeID, req.Title, req.Description, req.File,
	).Scan(&documentID)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetDocumentByID(documentID)
}

func (r *documentRepository) UpdateDocument(id int, req models.DocumentRequest) (*models.Document, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var typeID int

	if req.TypeID != 0 {
		typeID = req.TypeID
	} else if req.TypeName != "" {

		err = tx.QueryRow(`
			SELECT type_id FROM document_types WHERE type_name = $1
		`, req.TypeName).Scan(&typeID)

		if err == sql.ErrNoRows {

			err = tx.QueryRow(`
				INSERT INTO document_types (type_name) VALUES ($1) RETURNING type_id
			`, req.TypeName).Scan(&typeID)

			if err != nil {
				return nil, err
			}

		} else if err != nil {
			return nil, err
		}
	}

	var updatedID int

	err = tx.QueryRow(`
		UPDATE document
		SET type_id = $1, title = $2, description = $3, file = $4
		WHERE document_id = $5
		RETURNING document_id
	`,
		typeID, req.Title, req.Description, req.File, id,
	).Scan(&updatedID)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetDocumentByID(updatedID)
}

func (r *documentRepository) DeleteDocument(id int) error {
	result, err := r.db.Exec(
		"DELETE FROM document WHERE document_id = $1",
		id,
	)
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
