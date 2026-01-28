package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewUserStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserStorage {
	return &UserStorage{
		logger: logger,
		conn:   conn,
	}
}

func (s *UserStorage) CreateUser(userDto dto.CreateUser) (*int64, error) {
	var id int64
	query := `
		INSERT INTO users (email, password) 
		VALUES ($1, $2) 
		RETURNING id
	`

	err := s.conn.QueryRow(
		context.Background(),
		query,
		userDto.Email,
		userDto.PasswordHash,
	).Scan(&id)
	if err != nil {
		s.logger.Error("insert data to users table error", zap.Error(err))
		return nil, err
	}

	return &id, nil
}
