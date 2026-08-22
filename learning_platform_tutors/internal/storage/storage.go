package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	TutorReviewStorage  *TutorReviewStorage
	TutorOfferStorage   *TutorOfferStorage
	TutorStudentStorage *TutorStudentStorage
	TutorSubjectStorage *TutorSubjectStorage
}

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{
		TutorReviewStorage:  NewTutorReviewStorage(conn),
		TutorOfferStorage:   NewTutorOfferStorage(conn),
		TutorStudentStorage: NewTutorStudentStorage(conn),
		TutorSubjectStorage: NewTutorSubjectStorage(conn),
	}
}

func Connection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		return nil, fmt.Errorf("unable connect to postgres: %w", err)
	}
	return dbConn, nil
}
