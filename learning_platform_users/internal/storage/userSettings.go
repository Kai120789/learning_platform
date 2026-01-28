package storage

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
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

func (s *UserSettingsStorage) GetUserSettings(userId int64) (*models.UserSettings, error) {
	var userSettings models.UserSettings
	query := `
		SELECT *
		FROM user_settings
		WHERE user_id = $1
	`

	row := s.conn.QueryRow(context.Background(), query, userId)
	err := row.Scan(
		&userSettings.Id,
		&userSettings.UserId,
		&userSettings.Is2FaEnabled,
		&userSettings.IsNotificationsEnabled,
	)

	if err != nil {
		s.logger.Error("get user settings from db error", zap.Error(err))
		return nil, err
	}
	return &userSettings, nil
}

func (s *UserSettingsStorage) UpdateUserSettings(userSettings dto.UserSettings) error {
	query := `
		UPDATE user_settings
		SET
		    is_2fa_enabled = $2,
			is_notifications_enabled = $3
		WHERE user_id = $1
	`

	_, err := s.conn.Exec(
		context.Background(),
		query,
		userSettings.UserId,
		userSettings.Is2FaEnabled,
		userSettings.IsNotificationsEnabled,
	)

	if err != nil {
		s.logger.Error("update user settings in db error", zap.Error(err))
		return err
	}
	return nil
}
