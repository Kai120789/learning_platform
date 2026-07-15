package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/models/enum"
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

func (us *UserSettingsStorage) CreateUserSettings(userID int64, language enum.UserLanguage) error {
	query := `
		INSERT INTO user_settings (user_id, language)
		VALUES ($1, $2)
	`

	_, err := us.conn.Exec(
		context.Background(),
		query,
		userID,
		language,
	)

	if err != nil {
		return fmt.Errorf("insert settings for user %d: %w", userID, err)
	}

	return nil
}

func (us *UserSettingsStorage) GetUserSettings(userID int64) (*models.UserSettings, error) {
	var userSettings models.UserSettings
	query := `
		SELECT user_id, is_2fa_enabled, is_notifications_enabled, language, theme
		FROM user_settings
		WHERE user_id = $1
	`

	row := us.conn.QueryRow(context.Background(), query, userID)
	err := row.Scan(
		&userSettings.UserID,
		&userSettings.Is2FaEnabled,
		&userSettings.IsNotificationsEnabled,
		&userSettings.Language,
		&userSettings.Theme,
	)

	if err != nil {
		return nil, fmt.Errorf("get settings for user %d: %w", userID, err)
	}
	return &userSettings, nil
}

func (us *UserSettingsStorage) UpdateUserSettings(userSettings dto.UserSettingsRequest) error {
	query := `
		UPDATE user_settings
		SET
		    is_2fa_enabled = $2,
			is_notifications_enabled = $3,
			language = $4
		WHERE user_id = $1
	`

	_, err := us.conn.Exec(
		context.Background(),
		query,
		userSettings.UserID,
		userSettings.Is2FaEnabled,
		userSettings.IsNotificationsEnabled,
		userSettings.Language,
	)

	if err != nil {
		return fmt.Errorf("update info for user %d: %w", userSettings.UserID, err)
	}
	return nil
}

func (us *UserSettingsStorage) UpdateUserTheme(userID int64, theme enum.UserTheme) error {
	query := `
		UPDATE user_settings
		SET theme = $2
		WHERE user_id = $1
	`

	_, err := us.conn.Exec(context.Background(), query, userID, string(theme))
	if err != nil {
		return fmt.Errorf("update theme for user %d: %w", userID, err)
	}

	return nil
}
