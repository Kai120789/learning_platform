package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
)

type UserInfoStorage struct {
	conn *pgxpool.Pool
}

func NewUserInfoStorage(
	conn *pgxpool.Pool,
) *UserInfoStorage {
	return &UserInfoStorage{
		conn: conn,
	}
}

func (ui *UserInfoStorage) CreateUserInfo(userID int64, userDto dto.CreateUser) error {
	query := `
		INSERT INTO user_info (user_id, name, surname, patronymic, gender, birth_date) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := ui.conn.Exec(
		context.Background(),
		query,
		userID,
		userDto.Name,
		userDto.Surname,
		userDto.Patronymic,
		userDto.Gender,
		userDto.BirthDate,
	)

	if err != nil {
		return fmt.Errorf("insert info for user %d: %w", userID, err)
	}

	return nil
}

func (ui *UserInfoStorage) GetUserInfo(userID int64) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	query := `
		SElECT user_id, name, surname, patronymic, tg_username, city, about, avatar, gender, birth_date
		FROM user_info
		WHERE user_id = $1
	`

	row := ui.conn.QueryRow(context.Background(), query, userID)
	err := row.Scan(
		&userInfo.UserID,
		&userInfo.Name,
		&userInfo.Surname,
		&userInfo.Patronymic,
		&userInfo.TgUsername,
		&userInfo.City,
		&userInfo.About,
		&userInfo.Avatar,
		&userInfo.Gender,
		&userInfo.BirthDate,
	)

	if err != nil {
		return nil, fmt.Errorf("get info for user %d: %w", userID, err)
	}
	return &userInfo, nil
}

func (ui *UserInfoStorage) UpdateUserInfo(userInfo dto.UserInfoRequest) error {
	query := `
		UPDATE user_info
		SET
		    name = $2,
			surname = $3,
			patronymic = $4,
			city = $5,
			about = $6,
			gender = $7,
			birth_date = $8
		WHERE user_id = $1
	`

	_, err := ui.conn.Exec(
		context.Background(),
		query,
		userInfo.UserID,
		userInfo.Name,
		userInfo.Surname,
		userInfo.Patronymic,
		userInfo.City,
		userInfo.About,
		userInfo.Gender,
		userInfo.BirthDate,
	)

	if err != nil {
		return fmt.Errorf("update info for user %d: %w", userInfo.UserID, err)
	}
	return nil
}

func (ui *UserInfoStorage) UpdateUserAvatar(userID int64, avatar string) error {
	query := `
		UPDATE user_info
		SET avatar = $2
		WHERE user_id = $1
	`

	_, err := ui.conn.Exec(context.Background(), query, userID, avatar)
	if err != nil {
		return fmt.Errorf("update avatar for user %d: %w", userID, err)
	}

	return nil
}

func (ui *UserInfoStorage) GetUsersShortInfo(userIDs []int64) ([]dto.UserShortInfo, error) {
	var users []dto.UserShortInfo
	query := `
		SELECT user_id, name, surname, patronymic, tg_username
		FROM user_info
		WHERE user_id = ANY($1)
	`

	rows, err := ui.conn.Query(context.Background(), query, userIDs)
	if err != nil {
		return nil, fmt.Errorf("get users short info: %w", err)
	}

	for rows.Next() {
		var user dto.UserShortInfo

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.TgUsername,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one user short info: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (ui *UserInfoStorage) UpdateUserTgUsername(userID int64, tgUsername string) error {
	query := `
		UPDATE user_info
		SET tg_username = $2
		WHERE user_id = $1
	`

	_, err := ui.conn.Exec(context.Background(), query, userID, tgUsername)
	if err != nil {
		return fmt.Errorf("update tg username for user %d: %w", userID, err)
	}

	return nil
}
