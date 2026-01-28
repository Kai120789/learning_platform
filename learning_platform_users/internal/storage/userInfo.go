package storage

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
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
		s.logger.Error("get user info from db error", zap.Error(err))
		return nil, err
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
		s.logger.Error("update user info in db error", zap.Error(err))
		return err
	}
	return nil
}
