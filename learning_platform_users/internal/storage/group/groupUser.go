package group

import (
	"context"
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
)

type GroupUserStorage struct {
	conn *pgxpool.Pool
}

func NewGroupUserStorage(
	conn *pgxpool.Pool,
) *GroupUserStorage {
	return &GroupUserStorage{
		conn: conn,
	}
}

func (g *GroupUserStorage) AddUsersToGroup(userIds []int64, groupId int64) error {
	query := `
		INSERT INTO user_groups (user_id, group_id)
		SELECT unnest($1::bigint[]), $2
	`

	_, err := g.conn.Exec(context.Background(), query, userIds, groupId)
	if err != nil {
		return fmt.Errorf("add users %d to group %d: %w", userIds, groupId, err)
	}

	return nil
}

func (g *GroupUserStorage) RemoveUserFromGroup(userId int64, groupId int64) error {
	query := `DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2`

	_, err := g.conn.Exec(context.Background(), query, userId, groupId)
	if err != nil {
		return fmt.Errorf("remove user %d from group %d: %w", userId, groupId, err)
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
		return nil, fmt.Errorf("get all user %d groups from db: %w", userId, err)
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
			return nil, fmt.Errorf("scan one user %d group from db: %w", userId, err)
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
		return nil, fmt.Errorf("get all tutor %d groups from db: %w", tutorId, err)
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
			return nil, fmt.Errorf("scan one tutor %d group from db: %w", tutorId, err)
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
		return nil, fmt.Errorf("get all group %d users from db: %w", groupId, err)
	}

	for rows.Next() {
		var oneUser dto.ShortUserInfo

		err := rows.Scan(
			&oneUser.Id,
			&oneUser.Email,
			&oneUser.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one group %d user from db: %w", groupId, err)
		}

		users = append(users, oneUser)
	}

	return users, nil
}
