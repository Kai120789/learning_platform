package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/subjects/internal/models"
)

type UserSubjectStorage struct {
	conn *pgxpool.Pool
}

func NewUserSubjectStorage(conn *pgxpool.Pool) *UserSubjectStorage {
	return &UserSubjectStorage{
		conn: conn,
	}
}

func (u *UserSubjectStorage) GetAllUserSubjects(userID int64) ([]models.Subject, error) {
	var resSubjects []models.Subject

	query := `
		SELECT s.id, s.code, s.title, s.type
		FROM subjects AS s
		INNER JOIN user_subjects AS us 
		ON us.subject_id = s.id
		WHERE us.user_id = $1
	`

	rows, err := u.conn.Query(
		context.Background(),
		query,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("get all user %d subjects: %w", userID, err)
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

func (u *UserSubjectStorage) SetUserSubjects(userID int64, subjectIDs []int64) error {
	query := `
		INSERT INTO user_subjects (user_id, subject_id)
		SELECT $1, unnest($2::bigint[])
	`

	_, err := u.conn.Exec(
		context.Background(),
		query,
		userID,
		subjectIDs,
	)
	if err != nil {
		return fmt.Errorf("failed set user subjects: %w", err)
	}

	return nil
}

func (u *UserSubjectStorage) DeleteUserSubjects(userID int64, deletedSubjectIDs []int64) error {
	query := `
		DELETE FROM user_subjects
		WHERE subject_id = ANY($1::bigint[])
		AND user_id = $2
	`

	_, err := u.conn.Exec(context.Background(), query, deletedSubjectIDs, userID)
	if err != nil {
		return fmt.Errorf("delete subjects %d for user %d: %w", deletedSubjectIDs, userID, err)
	}

	return nil
}
