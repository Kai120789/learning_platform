package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	ScheduleStorage      *ScheduleStorage
	ScheduleSlotsStorage *ScheduleSlotsStorage
}

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{
		ScheduleStorage:      NewScheduleStorage(conn),
		ScheduleSlotsStorage: NewScheduleSlotsStorage(conn),
	}
}

func Connection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		return nil, fmt.Errorf("unable connect to postgres: %w", err)
	}

	return dbConn, err
}
