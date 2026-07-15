package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/models/enum"
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

func (u *UserBaseStorage) CreateUser(userDto dto.CreateUser) (*int64, error) {
	var id int64
	query := `
		INSERT INTO users (email, password_hash, role, status) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`

	status := enum.StatusActive
	if userDto.Role == enum.RoleTutor {
		status = enum.StatusInactive
	}

	err := u.conn.QueryRow(
		context.Background(),
		query,
		userDto.Email,
		userDto.PasswordHash,
		userDto.Role,
		status,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("insert user %s to db: %w", userDto.Email, err)
	}

	return &id, nil
}

func (u *UserBaseStorage) GetUserById(userID int64) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, email, password_hash, role, status
		FROM users
		WHERE id = $1
	`

	row := u.conn.QueryRow(context.Background(), query, userID)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("get user by id %d from db: %w", userID, err)
	}

	return &user, nil
}

func (u *UserBaseStorage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, email, password_hash, role, status
		FROM users
		WHERE email = $1
	`

	row := u.conn.QueryRow(context.Background(), query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("get user by email %s from db: %w", email, err)
	}

	return &user, nil
}

func (u *UserBaseStorage) ChangePassword(userID int64, newPasswordHash string) error {
	query := `
		UPDATE users
		SET password = $2
		WHERE id = $1
	`

	_, err := u.conn.Exec(context.Background(), query, userID, newPasswordHash)
	if err != nil {
		return fmt.Errorf("change password for user %d: %w", userID, err)
	}

	return nil
}

func (u *UserBaseStorage) ChangeEmail(userID int64, newEmail string) error {
	query := `
		UPDATE users
		SET email = $2
		WHERE id = $1
	`

	_, err := u.conn.Exec(context.Background(), query, userID, newEmail)
	if err != nil {
		return fmt.Errorf("change email for user %d: %w", userID, err)
	}

	return nil
}
