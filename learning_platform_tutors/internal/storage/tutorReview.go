package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorReviewStorage struct {
	conn *pgxpool.Pool
}

func NewTutorReviewStorage(conn *pgxpool.Pool) *TutorReviewStorage {
	return &TutorReviewStorage{
		conn: conn,
	}
}

func (tr *TutorReviewStorage) GetTutorReviews(
	getTutorReviews dto.GetTutorReviews,
) ([]models.TutorReview, int64, error) {
	var resReviews []models.TutorReview
	query := `
		SELECT
			id,
			tutor_id,
			author_id,
			subject_id,
			text,
			rating,
			created_at,
    		updated_at
		FROM tutor_reviews
		WHERE tutor_id = $1
		ORDER BY created_at DESC, id DESC
		LIMIT $2
		OFFSET $3
	`

	offset := (getTutorReviews.Page - 1) * getTutorReviews.Limit

	rows, err := tr.conn.Query(
		context.Background(),
		query,
		getTutorReviews.TutorID,
		getTutorReviews.Limit,
		offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("get tutor reviews: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneReview models.TutorReview
		err := rows.Scan(
			&oneReview.ID,
			&oneReview.TutorID,
			&oneReview.AuthorID,
			&oneReview.SubjectID,
			&oneReview.Text,
			&oneReview.Rating,
			&oneReview.CreatedAt,
			&oneReview.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan one review: %w", err)
		}

		resReviews = append(resReviews, oneReview)
	}

	count, err := tr.GetOneTutorReviewsCount(getTutorReviews.TutorID)
	if err != nil {
		return nil, 0, fmt.Errorf("get tutor reviews (get count): %w", err)
	}

	return resReviews, count, nil
}

func (tr *TutorReviewStorage) AddTutorReview(review dto.NewTutorReview) (*models.TutorReview, error) {
	var resReview models.TutorReview
	query := `
		INSERT INTO tutor_reviews (text, tutor_id, author_id, subject_id, rating)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, tutor_id, author_id, subject_id, text, rating, created_at, updated_at
	`

	err := tr.conn.QueryRow(
		context.Background(),
		query,
		review.Text,
		review.TutorID,
		review.AuthorID,
		review.SubjectID,
		review.Rating,
	).Scan(
		&resReview.ID,
		&resReview.TutorID,
		&resReview.AuthorID,
		&resReview.SubjectID,
		&resReview.Text,
		&resReview.Rating,
		&resReview.CreatedAt,
		&resReview.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("add tutor review: %w", err)
	}

	return &resReview, nil
}

func (tr *TutorReviewStorage) UpdateTutorReview(updReview dto.UpdateTutorReview) (*models.TutorReview, error) {
	var resReview models.TutorReview
	query := `
		UPDATE tutor_reviews
		SET 
		    text = $3,
		    subject_id = $4,
		    rating = $5,
		    updated_at = now()
		WHERE id = $1 AND author_id = $2
		RETURNING id, tutor_id, author_id, subject_id, text, rating, created_at, updated_at
	`

	err := tr.conn.QueryRow(
		context.Background(),
		query,
		updReview.ID,
		updReview.AuthorID,
		updReview.Text,
		updReview.SubjectID,
		updReview.Rating,
	).Scan(
		&resReview.ID,
		&resReview.TutorID,
		&resReview.AuthorID,
		&resReview.SubjectID,
		&resReview.Text,
		&resReview.Rating,
		&resReview.CreatedAt,
		&resReview.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("update tutor review: %w", err)
	}

	return &resReview, nil
}

func (tr *TutorReviewStorage) DeleteTutorReview(reviewID, authorID int64) error {
	query := `
		DELETE FROM tutor_reviews
		WHERE id = $1 AND author_id = $2
	`

	_, err := tr.conn.Exec(
		context.Background(),
		query,
		reviewID,
		authorID,
	)
	if err != nil {
		return fmt.Errorf("delete tutor review: %w", err)
	}

	return nil
}

func (tr *TutorReviewStorage) ChackCanAddReview(tutorID, authorID int64) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM tutor_reviews
			WHERE tutor_id = $1 AND author_id = $2
		)
	`

	err := tr.conn.QueryRow(
		context.Background(),
		query,
		tutorID,
		authorID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check can add review: %w", err)
	}

	return !exists, nil
}

func (tr *TutorReviewStorage) GetOneTutorReviewsCount(tutorID int64) (int64, error) {
	var count int64
	countQuery := `
		SELECT COUNT(*)
		FROM tutor_reviews
		WHERE tutor_id = $1
	`

	err := tr.conn.QueryRow(
		context.Background(),
		countQuery,
		tutorID,
	).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get count: %w", err)
	}

	return count, nil
}

func (tr *TutorReviewStorage) GetTutorsReviewsCount(tutorIDs []int64) ([]int64, error) {
	query := `
		SELECT
			t.tutor_id,
			COUNT(r.id) AS reviews_count
		FROM unnest($1::bigint[]) WITH ORDINALITY AS t(tutor_id, ord)
		LEFT JOIN tutor_reviews r ON r.tutor_id = t.tutor_id
		GROUP BY t.tutor_id, t.ord
		ORDER BY t.ord
	`

	rows, err := tr.conn.Query(
		context.Background(),
		query,
		tutorIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutors reviews count: %w", err)
	}
	defer rows.Close()

	result := make([]int64, 0, len(tutorIDs))

	for rows.Next() {
		var tutorID int64
		var count int64

		if err := rows.Scan(&tutorID, &count); err != nil {
			return nil, fmt.Errorf("scan tutor reviews count: %w", err)
		}

		result = append(result, count)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tutors reviews count: %w", err)
	}

	return result, nil
}

func (tr *TutorReviewStorage) GetReviewByTutorAndAuthor(tutorID, userID int64) (*models.TutorReview, error) {
	var resReview models.TutorReview
	query := `
		SELECT
			id,
			tutor_id,
			author_id,
			subject_id,
			text,
			rating,
			created_at,
    		updated_at
		FROM tutor_reviews
		WHERE tutor_id = $1 AND author_id = $2
	`

	err := tr.conn.QueryRow(
		context.Background(),
		query,
		tutorID,
		userID,
	).Scan(
		&resReview.ID,
		&resReview.TutorID,
		&resReview.AuthorID,
		&resReview.SubjectID,
		&resReview.Text,
		&resReview.Rating,
		&resReview.CreatedAt,
		&resReview.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get review by tutor and author: %w", err)
	}

	return &resReview, nil
}

func (tr *TutorReviewStorage) GetOneTutorRating(tutorID int64) (float32, error) {
	var rating float32
	query := `
		SELECT COALESCE(AVG(rating), 0)
		FROM tutor_reviews
		WHERE tutor_id = $1
	`

	err := tr.conn.QueryRow(
		context.Background(),
		query,
		tutorID,
	).Scan(&rating)
	if err != nil {
		return 0, fmt.Errorf("get tutor rating: %w", err)
	}

	return rating, nil
}

func (tr *TutorReviewStorage) GetTutorsRatings(tutorIDs []int64) ([]float32, error) {
	query := `
		SELECT
			t.tutor_id,
			COALESCE(AVG(r.rating), 0)::float8 AS rating
		FROM unnest($1::bigint[]) WITH ORDINALITY AS t(tutor_id, ord)
		LEFT JOIN tutor_reviews r ON r.tutor_id = t.tutor_id
		GROUP BY t.tutor_id, t.ord
		ORDER BY t.ord
	`

	rows, err := tr.conn.Query(
		context.Background(),
		query,
		tutorIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutors ratings: %w", err)
	}
	defer rows.Close()

	result := make([]float32, 0, len(tutorIDs))

	for rows.Next() {
		var tutorID int64
		var rating float32

		if err := rows.Scan(&tutorID, &rating); err != nil {
			return nil, fmt.Errorf("scan tutor rating: %w", err)
		}

		result = append(result, rating)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tutors ratings: %w", err)
	}

	return result, nil
}
