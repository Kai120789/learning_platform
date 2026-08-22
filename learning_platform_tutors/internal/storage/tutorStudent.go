package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorStudentStorage struct {
	conn *pgxpool.Pool
}

func NewTutorStudentStorage(conn *pgxpool.Pool) *TutorStudentStorage {
	return &TutorStudentStorage{
		conn: conn,
	}
}

func (ts *TutorStudentStorage) AddStudent(tutorID, studentID int64) error {
	query := `
		INSERT INTO tutor_students (tutor_id, student_id)
		VALUES ($1, $2)
		ON CONFLICT (tutor_id, student_id) DO NOTHING
	`

	_, err := ts.conn.Exec(context.Background(), query, tutorID, studentID)
	if err != nil {
		return fmt.Errorf("add tutor student: %w", err)
	}

	return nil
}

func (ts *TutorStudentStorage) DeleteOneStudent(tutorID, studentID int64) error {
	query := `
		DELETE FROM tutor_students
		WHERE tutor_id = $1 AND student_id = $2
	`

	_, err := ts.conn.Exec(context.Background(), query, tutorID, studentID)
	if err != nil {
		return fmt.Errorf("delete one tutor student: %w", err)
	}

	return nil
}

func (ts *TutorStudentStorage) DeleteStudents(tutorID int64, studentIDs []int64) error {
	query := `
		DELETE FROM tutor_students
		WHERE tutor_id = $1 AND student_id = ANY($2)
	`

	_, err := ts.conn.Exec(context.Background(), query, tutorID, studentIDs)
	if err != nil {
		return fmt.Errorf("delete tutor students: %w", err)
	}

	return nil
}

func (ts *TutorStudentStorage) GetTutorStudents(
	getTutorStudents dto.GetTutorStudents,
) ([]models.TutorStudent, int64, error) {
	var tutorStudents []models.TutorStudent
	query := `
		SELECT
			id,
			tutor_id,
			student_id,
			last_interacted_at
		FROM tutor_students
		WHERE tutor_id = $1 AND (
		  	$2::BIGINT IS NULL
		  	OR last_interacted_at >= now() - ($2 * INTERVAL '1 day')
	  	)
		ORDER BY last_interacted_at DESC NULLS LAST, id DESC
		LIMIT $3
		OFFSET $4
	`

	offset := (getTutorStudents.Page - 1) * getTutorStudents.Limit

	rows, err := ts.conn.Query(
		context.Background(),
		query,
		getTutorStudents.TutorID,
		getTutorStudents.InteractedWithinDays,
		getTutorStudents.Limit,
		offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("get tutor students: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneTutorStudent models.TutorStudent
		err := rows.Scan(
			&oneTutorStudent.ID,
			&oneTutorStudent.TutorID,
			&oneTutorStudent.StudentID,
			&oneTutorStudent.LastInteractedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan one tutor student: %w", err)
		}

		tutorStudents = append(tutorStudents, oneTutorStudent)
	}

	count, err := ts.GetOneTutorStudentsCount(getTutorStudents.TutorID, getTutorStudents.InteractedWithinDays)
	if err != nil {
		return nil, 0, fmt.Errorf("get students (get count): %w", err)
	}

	return tutorStudents, count, err
}

func (ts *TutorStudentStorage) UpdateLastInteracted(tutorID, studentID int64) error {
	query := `
		UPDATE tutor_students
		SET
			last_interacted_at = now(),
			updated_at = now()
		WHERE tutor_id = $1 AND student_id = $2
	`

	_, err := ts.conn.Exec(
		context.Background(),
		query,
		tutorID,
		studentID,
	)
	if err != nil {
		return fmt.Errorf("update last interacted: %w", err)
	}

	return nil
}

func (ts *TutorStudentStorage) GetOneTutorStudentsCount(tutorID int64, interactedWithinDays *int64) (int64, error) {
	var count int64
	countQuery := `
		SELECT COUNT(*)
		FROM tutor_students
		WHERE tutor_id = $1 AND (
	      $2::bigint IS NULL
	      OR last_interacted_at >= now() - ($2 * INTERVAL '1 day')
	  	)
	`

	err := ts.conn.QueryRow(
		context.Background(),
		countQuery,
		tutorID,
		interactedWithinDays,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get one tutor students count: %w", err)
	}

	return count, nil
}

func (ts *TutorStudentStorage) GetTutorsStudentsCount(tutorIDs []int64) ([]int64, error) {
	query := `
		SELECT
			t.tutor_id,
			COUNT(ts.id) AS students_count
		FROM unnest($1::bigint[]) WITH ORDINALITY AS t(tutor_id, ord)
		LEFT JOIN tutor_students ts ON ts.tutor_id = t.tutor_id
		GROUP BY t.tutor_id, t.ord
		ORDER BY t.ord
	`

	rows, err := ts.conn.Query(
		context.Background(),
		query,
		tutorIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutors students count: %w", err)
	}
	defer rows.Close()

	counts := make([]int64, 0, len(tutorIDs))
	for rows.Next() {
		var tutorID int64
		var count int64

		if err := rows.Scan(&tutorID, &count); err != nil {
			return nil, fmt.Errorf("scan tutor students count: %w", err)
		}

		counts = append(counts, count)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tutors students count: %w", err)
	}

	return counts, nil
}
