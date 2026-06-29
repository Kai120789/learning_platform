package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	SubjectStorage     *SubjectStorage
	UserSubjectStorage *UserSubjectStorage
}

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{
		SubjectStorage:     NewSubjectStorage(conn),
		UserSubjectStorage: NewUserSubjectStorage(conn),
	}
}

func Connection(connStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable connect to postgres: %w", err)
	}

	return dbConn, nil
}
