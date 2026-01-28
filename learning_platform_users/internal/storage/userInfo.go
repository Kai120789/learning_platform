package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserInfoStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewUserInfoStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserInfoStorage {
	return &UserInfoStorage{
		logger: logger,
		conn:   conn,
	}
}

func (s *UserInfoStorage) CreateUserInfo(userId int64, userDto dto.CreateUser) error {
	query := `
		INSERT INTO user_info (user_id, name, surname, lastname, role) 
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.conn.Exec(
		context.Background(),
		query,
		userId,
		userDto.Name,
		userDto.Surname,
		userDto.LastName,
		userDto.Role,
	)

	if err != nil {
		s.logger.Error("insert data to user_info table error", zap.Error(err))
		return err
	}

	return nil
}
