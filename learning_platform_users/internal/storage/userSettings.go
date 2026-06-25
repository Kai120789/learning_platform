package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
)

type UserSettingsStorage struct {
	conn *pgxpool.Pool
}

func NewUserSettingsStorage(
	conn *pgxpool.Pool,
) *UserSettingsStorage {
	return &UserSettingsStorage{
		conn: conn,
	}
}

func (s *UserSettingsStorage) CreateUserSettings(userID int64) error {
	query := `
		INSERT INTO user_settings (user_id)
		VALUES ($1)
	`

	_, err := s.conn.Exec(
		context.Background(),
		query, userID,
	)

	if err != nil {
		return fmt.Errorf("insert settings for user %d: %w", userID, err)
	}

	return nil
}

func (s *UserSettingsStorage) GetUserSettings(userID int64) (*models.UserSettings, error) {
	var userSettings models.UserSettings
	query := `
		SELECT user_id, is_2fa_enabled, is_notifications_enabled
		FROM user_settings
		WHERE user_id = $1
	`

	row := s.conn.QueryRow(context.Background(), query, userID)
	err := row.Scan(
		&userSettings.UserID,
		&userSettings.Is2FaEnabled,
		&userSettings.IsNotificationsEnabled,
	)

	if err != nil {
		return nil, fmt.Errorf("get settings for user %d: %w", userID, err)
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
		userSettings.UserID,
		userSettings.Is2FaEnabled,
		userSettings.IsNotificationsEnabled,
	)

	if err != nil {
		return fmt.Errorf("update info for user %d: %w", userSettings.UserID, err)
	}
	return nil
}
