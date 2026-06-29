package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	LessonStorage      *LessonStorage
	LessonMediaStorage *LessonMediaStorage
	LessonUserStorage  *LessonUserStorage
}

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{
		LessonStorage:      NewLessonStorage(conn),
		LessonMediaStorage: NewLessonMediaStorage(conn),
		LessonUserStorage:  NewLessonUserStorage(conn),
	}
}

func Connection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		return nil, fmt.Errorf("unable connect to postgres: %w", err)
	}

	return dbConn, err
}
