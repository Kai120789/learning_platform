package user

import (
	"context"
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
)

type UserBaseStorage struct {
	conn *pgxpool.Pool
}

func NewUserBaseStorage(
	conn *pgxpool.Pool,
) *UserBaseStorage {
	return &UserBaseStorage{
		conn: conn,
	}
}

func (s *UserBaseStorage) CreateUser(userDto dto.CreateUser) (*int64, error) {
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
		return nil, fmt.Errorf("insert user %s to db: %w", userDto.Email, err)
	}

	return &id, nil
}

func (s *UserBaseStorage) GetUserById(userId int64) (*models.User, error) {
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
		return nil, fmt.Errorf("get user by id %d from db: %w", userId, err)
	}

	return &user, nil
}

func (s *UserBaseStorage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT *
		FROM users
		WHERE email = $1
	`

	row := s.conn.QueryRow(context.Background(), query, email)

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, fmt.Errorf("get user by email %s from db: %w", email, err)
	}

	return &user, nil
}

func (s *UserBaseStorage) ChangePassword(userId int64, newPasswordHash string) error {
	query := `
		UPDATE users
		SET password = $2
		WHERE id = $1
	`

	_, err := s.conn.Exec(context.Background(), query, userId, newPasswordHash)
	if err != nil {
		return fmt.Errorf("change password for user %d: %w", userId, err)
	}

	return nil
}

func (s *UserBaseStorage) ChangeEmail(userId int64, newEmail string) error {
	query := `
		UPDATE users
		SET email = $2
		WHERE id = $1
	`

	_, err := s.conn.Exec(context.Background(), query, userId, newEmail)
	if err != nil {
		return fmt.Errorf("change email for user %d: %w", userId, err)
	}

	return nil
}
