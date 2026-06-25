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

func (s *UserInfoStorage) CreateUserInfo(userID int64, userDto dto.CreateUser) error {
	query := `
		INSERT INTO user_info (user_id, name, surname, lastname) 
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.conn.Exec(
		context.Background(),
		query,
		userID,
		userDto.Name,
		userDto.Surname,
		userDto.LastName,
	)

	if err != nil {
		return fmt.Errorf("insert info for user %d: %w", userID, err)
	}

	return nil
}

func (s *UserInfoStorage) GetUserInfo(userID int64) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	query := `
		SElECT user_id, name, surname, lastname, city, about
		FROM user_info
		WHERE user_id = $1
	`

	row := s.conn.QueryRow(context.Background(), query, userID)
	err := row.Scan(
		&userInfo.UserID,
		&userInfo.Name,
		&userInfo.Surname,
		&userInfo.Lastname,
		&userInfo.City,
		&userInfo.About,
	)

	if err != nil {
		return nil, fmt.Errorf("get info for user %d: %w", userID, err)
	}
	return &userInfo, nil
}

func (s *UserInfoStorage) UpdateUserInfo(userInfo dto.UserInfo) error {
	query := `
		UPDATE user_info
		SET
		    name = $2,
			surname = $3,
			lastname = $4,
			city = $5,
			about = $6,
		WHERE user_id = $1
	`

	_, err := s.conn.Exec(
		context.Background(),
		query,
		userInfo.UserID,
		userInfo.Name,
		userInfo.Surname,
		userInfo.Lastname,
		userInfo.City,
		userInfo.About,
	)

	if err != nil {
		return fmt.Errorf("update info for user %d: %w", userInfo.UserID, err)
	}
	return nil
}
