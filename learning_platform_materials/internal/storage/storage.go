package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	MaterialStorage       *MaterialStorage
	MaterialFolderStorage *MaterialFolderStorage
	MaterialUserStorage   *MaterialUserStorage
}

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{
		MaterialStorage:       NewMaterialStorage(conn),
		MaterialFolderStorage: NewMaterialFolderStorage(conn),
		MaterialUserStorage:   NewMaterialUserStorage(conn),
	}
}

func Connection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		return nil, fmt.Errorf("unable connect to postgres: %w", err)
	}
	return dbConn, nil
}
