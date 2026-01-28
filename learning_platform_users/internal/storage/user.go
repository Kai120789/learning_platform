package storage

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
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

func (s *UserStorage) GetUser(userId int64) (*models.User, error) {
	var user models.User
	query := `
		SELECT *
		FROM users
		WHERE id = $1
	`

	row := s.conn.QueryRow(context.Background(), query, userId)

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		s.logger.Error("get user from db error", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) ChangePassword(userId int64, newPasswordHash string) error {
	query := `
		UPDATE users
		SET password = $2
		WHERE id = $1
	`

	_, err := s.conn.Exec(context.Background(), query, userId, newPasswordHash)
	if err != nil {
		s.logger.Error("update password in users table error", zap.Error(err))
		return err
	}

	return nil
}

func (s *UserStorage) ChangeEmail(userId int64, newEmail string) error {
	query := `
		UPDATE users
		SET email = $2
		WHERE id = $1
	`

	_, err := s.conn.Exec(context.Background(), query, userId, newEmail)
	if err != nil {
		s.logger.Error("update email in users table error", zap.Error(err))
		return err
	}

	return nil
}
