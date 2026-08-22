package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/tutors/internal/dto"
)

type TutorSubjectStorage struct {
	conn *pgxpool.Pool
}

func NewTutorSubjectStorage(conn *pgxpool.Pool) *TutorSubjectStorage {
	return &TutorSubjectStorage{
		conn: conn,
	}
}

func (tsu *TutorSubjectStorage) AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	query := `
		INSERT INTO tutor_subjects (tutor_id, subject_id)
		SELECT $1, UNNEST($2::BIGINT[])
		ON CONFLICT (tutor_id, subject_id) DO NOTHING
		RETURNING subject_id
	`

	rows, err := tsu.conn.Query(
		context.Background(),
		query,
		tutorID,
		subjectIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("add tutor subjects: %w", err)
	}
	defer rows.Close()

	var addedSubjectIDs []int64
	for rows.Next() {
		var subjectID int64
		if err := rows.Scan(&subjectID); err != nil {
			return nil, fmt.Errorf("scan added subject id: %w", err)
		}

		addedSubjectIDs = append(addedSubjectIDs, subjectID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate added subjects: %w", err)
	}

	return addedSubjectIDs, nil
}

func (tsu *TutorSubjectStorage) UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	tx, err := tsu.conn.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("begin update tutor subjects: %w", err)
	}
	defer tx.Rollback(context.Background())

	deleteQuery := `
		DELETE FROM tutor_subjects
		WHERE tutor_id = $1 AND NOT (subject_id = ANY($2::BIGINT[]))
	`

	_, err = tx.Exec(
		context.Background(),
		deleteQuery,
		tutorID,
		subjectIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("delete old tutor subjects: %w", err)
	}

	addQuery := `
		INSERT INTO tutor_subjects (tutor_id, subject_id)
		SELECT $1, UNNEST($2::BIGINT[])
		ON CONFLICT (tutor_id, subject_id) DO NOTHING
		RETURNING subject_id
	`

	rows, err := tx.Query(
		context.Background(),
		addQuery,
		tutorID,
		subjectIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("add new tutor subjects: %w", err)
	}
	defer rows.Close()

	var resSubjectIDs []int64
	for rows.Next() {
		var subjectID int64
		if err := rows.Scan(&subjectID); err != nil {
			return nil, fmt.Errorf("scan added subject id: %w", err)
		}

		resSubjectIDs = append(resSubjectIDs, subjectID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate added subjects: %w", err)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, fmt.Errorf("commit update tutor subjects: %w", err)
	}

	return resSubjectIDs, nil
}

func (tsu *TutorSubjectStorage) GetOneTutorSubjectIDs(tutorID int64) ([]int64, error) {
	var resSubjectIDs []int64
	query := `
		SELECT subject_id 
		FROM tutor_subjects
		WHERE tutor_id = $1
	`

	rows, err := tsu.conn.Query(context.Background(), query, tutorID)
	if err != nil {
		return nil, fmt.Errorf("get tutor subjects: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var subjectID int64
		if err := rows.Scan(&subjectID); err != nil {
			return nil, fmt.Errorf("scan subject id: %w", err)
		}

		resSubjectIDs = append(resSubjectIDs, subjectID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate subjects: %w", err)
	}

	return resSubjectIDs, nil
}

func (tsu *TutorSubjectStorage) GetTutorsSubjectIDs(tutorIDs []int64) ([][]int64, error) {
	query := `
		SELECT
			t.tutor_id,
			COALESCE(
				ARRAY_AGG(ts.subject_id ORDER BY ts.subject_id)
					FILTER (WHERE ts.subject_id IS NOT NULL),
				'{}'
			) AS subject_ids
		FROM unnest($1::bigint[]) WITH ORDINALITY AS t(tutor_id, ord)
		LEFT JOIN tutor_subjects ts ON ts.tutor_id = t.tutor_id
		GROUP BY t.tutor_id, t.ord
		ORDER BY t.ord
	`

	rows, err := tsu.conn.Query(
		context.Background(),
		query,
		tutorIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutors subject ids: %w", err)
	}
	defer rows.Close()

	result := make([][]int64, 0, len(tutorIDs))

	for rows.Next() {
		var tutorID int64
		var subjectIDs []int64

		if err := rows.Scan(&tutorID, &subjectIDs); err != nil {
			return nil, fmt.Errorf("scan tutor subject ids: %w", err)
		}

		result = append(result, subjectIDs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tutors subject ids: %w", err)
	}

	return result, nil
}

func (tsu *TutorSubjectStorage) GetTutorsBySubject(getTutors dto.GetTutors) ([]int64, int64, error) {
	query := `
		SELECT DISTINCT tutor_id
		FROM tutor_subjects
		WHERE ($1::bigint IS NULL OR subject_id = $1)
		ORDER BY tutor_id
		LIMIT $2
		OFFSET $3
	`

	offset := (getTutors.Page - 1) * getTutors.Limit

	rows, err := tsu.conn.Query(
		context.Background(),
		query,
		getTutors.SubjectID,
		getTutors.Limit,
		offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("get tutors: %w", err)
	}
	defer rows.Close()

	tutorIDs := make([]int64, 0, getTutors.Limit)

	for rows.Next() {
		var tutorID int64

		if err := rows.Scan(&tutorID); err != nil {
			return nil, 0, fmt.Errorf("scan tutor id: %w", err)
		}

		tutorIDs = append(tutorIDs, tutorID)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterate tutors: %w", err)
	}

	count, err := tsu.GetTutorsCount(getTutors.SubjectID)
	if err != nil {
		return nil, 0, fmt.Errorf("get tutors (get tutors count): %w", err)
	}

	return tutorIDs, count, nil
}

func (tsu *TutorSubjectStorage) GetTutorsCount(subjectID *int64) (int64, error) {
	var count int64
	countQuery := `
		SELECT COUNT(DISTINCT tutor_id)
		FROM tutor_subjects
		WHERE ($1::bigint IS NULL OR subject_id = $1)
	`

	err := tsu.conn.QueryRow(
		context.Background(),
		countQuery,
		subjectID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get tutors count: %w", err)
	}

	return count, nil
}
