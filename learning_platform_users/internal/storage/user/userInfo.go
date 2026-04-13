package user

import (
	"context"
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
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

func (s *UserInfoStorage) CreateUserInfo(userId int64, userDto dto.CreateUser) error {
	query := `
		INSERT INTO user_info (user_id, name, surname, lastname, role, status) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	status := "ACTIVE"
	if userDto.Role == "TUTOR" {
		status = "INACTIVE"
	}

	_, err := s.conn.Exec(
		context.Background(),
		query,
		userId,
		userDto.Name,
		userDto.Surname,
		userDto.LastName,
		userDto.Role,
		status,
	)

	if err != nil {
		return fmt.Errorf("insert info for user %d: %w", userId, err)
	}

	return nil
}

func (s *UserInfoStorage) GetUserInfo(userId int64) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	query := `
		SElECT *
		FROM user_info
		WHERE user_id = $1
	`

	row := s.conn.QueryRow(context.Background(), query, userId)
	err := row.Scan(
		&userInfo.Id,
		&userInfo.UserId,
		&userInfo.Name,
		&userInfo.Surname,
		&userInfo.Lastname,
		&userInfo.City,
		&userInfo.About,
		&userInfo.Role,
		&userInfo.Status,
		&userInfo.Class,
	)

	if err != nil {
		return nil, fmt.Errorf("get info for user %d: %w", userId, err)
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
			class = $7
		WHERE user_id = $1
	`

	_, err := s.conn.Exec(
		context.Background(),
		query,
		userInfo.UserId,
		userInfo.Name,
		userInfo.Surname,
		userInfo.Lastname,
		userInfo.City,
		userInfo.About,
		userInfo.Class,
	)

	if err != nil {
		return fmt.Errorf("update info for user %d: %w", userInfo.UserId, err)
	}
	return nil
}
