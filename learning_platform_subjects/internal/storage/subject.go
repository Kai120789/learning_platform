package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/subjects/internal/models"
)

type SubjectStorage struct {
	conn *pgxpool.Pool
}

func NewSubjectStorage(conn *pgxpool.Pool) *SubjectStorage {
	return &SubjectStorage{
		conn: conn,
	}
}

func (s *SubjectStorage) GetOneSubject(subjectID int64) (*models.Subject, error) {
	var resSubject models.Subject
	query := `
		SELECT id, code, title, type
		FROM subjects
		WHERE id = $1
	`

	err := s.conn.QueryRow(
		context.Background(),
		query,
		subjectID,
	).Scan(
		&resSubject.ID,
		&resSubject.Code,
		&resSubject.Title,
		&resSubject.Type,
	)
	if err != nil {
		return nil, fmt.Errorf("get one subject: %w", err)
	}

	return &resSubject, nil
}

func (s *SubjectStorage) GetAllSubjects() ([]models.Subject, error) {
	var resSubjects []models.Subject

	query := `
		SELECT id, code, title, type
		FROM subjects
	`

	rows, err := s.conn.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("get all subjects: %w", err)
	}

	for rows.Next() {
		var oneSubject models.Subject

		err := rows.Scan(
			&oneSubject.ID,
			&oneSubject.Code,
			&oneSubject.Title,
			&oneSubject.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("failed scan one subject: %w", err)
		}

		resSubjects = append(resSubjects, oneSubject)
	}

	return resSubjects, nil
}
