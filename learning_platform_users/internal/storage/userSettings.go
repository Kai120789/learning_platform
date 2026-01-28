package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserSettingsStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewUserSettingsStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserSettingsStorage {
	return &UserSettingsStorage{
		logger: logger,
		conn:   conn,
	}
}

func (s *UserSettingsStorage) CreateUserSettings(userId int64) error {
	query := `
		INSERT INTO user_settings (user_id)
		VALUES ($1)
	`

	_, err := s.conn.Exec(
		context.Background(),
		query, userId,
	)

	if err != nil {
		s.logger.Error("insert data to user_settings table error", zap.Error(err))
		return err
	}

	return nil
}
