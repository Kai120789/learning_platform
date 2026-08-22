package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorOfferStorage struct {
	conn *pgxpool.Pool
}

func NewTutorOfferStorage(conn *pgxpool.Pool) *TutorOfferStorage {
	return &TutorOfferStorage{
		conn: conn,
	}
}

func (to *TutorOfferStorage) GetTutorOffers(tutorID int64) ([]models.TutorOffer, error) {
	var tutorOffers []models.TutorOffer
	query := `
		SELECT id, tutor_id, subject_id, title, description, price, duration_minutes
		FROM tutor_offers
		WHERE tutor_id = $1
		ORDER BY subject_id ASC, price ASC
	`

	rows, err := to.conn.Query(context.Background(), query, tutorID)
	if err != nil {
		return nil, fmt.Errorf("get tutor offers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneTutorOffer models.TutorOffer
		err := rows.Scan(
			&oneTutorOffer.ID,
			&oneTutorOffer.TutorID,
			&oneTutorOffer.SubjectID,
			&oneTutorOffer.Title,
			&oneTutorOffer.Description,
			&oneTutorOffer.Price,
			&oneTutorOffer.DurationMinutes,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one tutor offer: %w", err)
		}

		tutorOffers = append(tutorOffers, oneTutorOffer)
	}

	return tutorOffers, nil
}

func (to *TutorOfferStorage) AddTutorOffer(offer dto.NewTutorOffer) (*models.TutorOffer, error) {
	var resOffer models.TutorOffer
	query := `
		INSERT INTO tutor_offers (tutor_id, subject_id, title, description, price, duration_minutes)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, tutor_id, subject_id, title, description, price, duration_minutes
	`

	err := to.conn.QueryRow(
		context.Background(),
		query,
		offer.TutorID,
		offer.SubjectID,
		offer.Title,
		offer.Description,
		offer.Price,
		offer.DurationMinutes,
	).Scan(
		&resOffer.ID,
		&resOffer.TutorID,
		&resOffer.SubjectID,
		&resOffer.Title,
		&resOffer.Description,
		&resOffer.Price,
		&resOffer.DurationMinutes,
	)
	if err != nil {
		return nil, fmt.Errorf("add tutor offer: %w", err)
	}

	return &resOffer, nil
}

func (to *TutorOfferStorage) UpdateTutorOffer(updOffer dto.UpdateTutorOffer) (*models.TutorOffer, error) {
	var resOffer models.TutorOffer
	query := `
		UPDATE tutor_offers
		SET 
			subject_id = $3,
			title = $4,
			description = $5 ,
			price = $6,
			duration_minutes = $7,
			updated_at = now()
		WHERE id = $1 AND tutor_id = $2
		RETURNING id, tutor_id, subject_id, title, description, price, duration_minutes
	`

	err := to.conn.QueryRow(
		context.Background(),
		query,
		updOffer.ID,
		updOffer.TutorID,
		updOffer.SubjectID,
		updOffer.Title,
		updOffer.Description,
		updOffer.Price,
		updOffer.DurationMinutes,
	).Scan(
		&resOffer.ID,
		&resOffer.TutorID,
		&resOffer.SubjectID,
		&resOffer.Title,
		&resOffer.Description,
		&resOffer.Price,
		&resOffer.DurationMinutes,
	)
	if err != nil {
		return nil, fmt.Errorf("update tutor offer: %w", err)
	}

	return &resOffer, nil
}

func (to *TutorOfferStorage) DeleteOneTutorOffer(offerID, tutorID int64) error {
	query := `
		DELETE FROM tutor_offers
		WHERE id = $1 AND tutor_id = $2
	`

	_, err := to.conn.Exec(context.Background(), query, offerID, tutorID)
	if err != nil {
		return fmt.Errorf("delete one tutor offer: %w", err)
	}

	return nil
}

func (to *TutorOfferStorage) DeleteTutorOffers(offerIDs []int64, tutorID int64) error {
	query := `
		DELETE FROM tutor_offers
		WHERE id = ANY($1) AND tutor_id = $2
	`

	_, err := to.conn.Exec(context.Background(), query, offerIDs, tutorID)
	if err != nil {
		return fmt.Errorf("delete tutor offers: %w", err)
	}

	return nil
}
