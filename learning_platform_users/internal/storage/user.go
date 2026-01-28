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

func (s *UserStorage) GetUser(userId int64) (*dto.GetUser, error) {
	var res dto.GetUser
	query := `
		SELECT *
		FROM users
		WHERE id = $1
	`

	row := s.conn.QueryRow(context.Background(), query, userId)

	err := row.Scan(
		&res.UserId,
		&res.Email,
		&res.PasswordHash,
	)

	if err != nil {
		s.logger.Error("get user from db error", zap.Error(err))
		return nil, err
	}

	return &res, nil
}
