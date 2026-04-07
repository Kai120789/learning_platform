package group

import (
	"context"
	"github.com/Kai120789/learning_platform_models/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type GroupUserStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewGroupUserStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *GroupUserStorage {
	return &GroupUserStorage{
		logger: logger,
		conn:   conn,
	}
}

func (g *GroupUserStorage) AddUsersToGroup(userId int64, groupId int64) error {
	query := `INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2)`

	_, err := g.conn.Exec(context.Background(), query, userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupUserStorage) RemoveUserFromGroup(userId int64, groupId int64) error {
	query := `DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2`

	_, err := g.conn.Exec(context.Background(), query, userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupUserStorage) GetUserGroups(userId int64) ([]models.Group, error) {
	var groups []models.Group

	query := `
		SELECT g.* FROM groups AS g
		JOIN user_groups AS ug ON ug.group_id = g.id AND ug.user_id = $1
	`

	rows, err := g.conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var oneGroup models.Group
		err := rows.Scan(
			&oneGroup.Id,
			&oneGroup.Title,
			&oneGroup.Description,
			&oneGroup.SubjectId,
			&oneGroup.TutorId,
			&oneGroup.TgGroupLink,
			&oneGroup.TgChatId,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, oneGroup)
	}

	return groups, nil
}

func (g *GroupUserStorage) GetGroupsByTutorId(tutorId int64) ([]models.Group, error) {
	var groups []models.Group

	query := `
		SELECT * FROM groups
		WHERE tutor_id = $1
	`

	rows, err := g.conn.Query(context.Background(), query, tutorId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var oneGroup models.Group
		err := rows.Scan(
			&oneGroup.Id,
			&oneGroup.Title,
			&oneGroup.Description,
			&oneGroup.SubjectId,
			&oneGroup.TutorId,
			&oneGroup.TgGroupLink,
			&oneGroup.TgChatId,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, oneGroup)
	}

	return groups, nil
}

func (g *GroupUserStorage) GetGroupUsers(groupId int64) ([]dto.ShortUserInfo, error) {
	var users []dto.ShortUserInfo

	query := `
		SELECT u.id, u.email, ui.name FROM users AS u
		JOIN user_info AS ui ON u.id = ui.user_id
		JOIN user_groups AS ug ON ug.group_id = $1 AND u.id = ug.user_id
	`

	rows, err := g.conn.Query(context.Background(), query, groupId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var oneUser dto.ShortUserInfo

		err := rows.Scan(
			&oneUser.Id,
			&oneUser.Email,
			&oneUser.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, oneUser)
	}

	return users, nil
}
