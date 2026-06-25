package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/groups/internal/models"
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

func (g *GroupUserStorage) AddUsersToGroup(userIDs []int64, groupID int64) error {
	query := `
		INSERT INTO user_groups (user_id, group_id)
		SELECT unnest($1::bigint[]), $2
	`

	_, err := g.conn.Exec(context.Background(), query, userIDs, groupID)
	if err != nil {
		return fmt.Errorf("add users %d to group %d: %w", userIDs, groupID, err)
	}

	return nil
}

func (g *GroupUserStorage) RemoveUserFromGroup(userID int64, groupID int64) error {
	query := `DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2`

	_, err := g.conn.Exec(context.Background(), query, userID, groupID)
	if err != nil {
		return fmt.Errorf("remove user %d from group %d: %w", userID, groupID, err)
	}

	return nil
}

func (g *GroupUserStorage) GetUserGroups(userID int64) ([]models.Group, error) {
	var groups []models.Group

	query := `
		SELECT g.id, g.title, g.description, g.subject_id, g.tutor_id, g.tg_group_link, g.tg_chat_id
		FROM groups AS g
		JOIN user_groups AS ug ON ug.group_id = g.id AND ug.user_id = $1
	`

	rows, err := g.conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("get all user %d groups from db: %w", userID, err)
	}

	for rows.Next() {
		var oneGroup models.Group
		err := rows.Scan(
			&oneGroup.ID,
			&oneGroup.Title,
			&oneGroup.Description,
			&oneGroup.SubjectID,
			&oneGroup.TutorID,
			&oneGroup.TgGroupLink,
			&oneGroup.TgChatID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one user %d group from db: %w", userID, err)
		}

		groups = append(groups, oneGroup)
	}

	return groups, nil
}

func (g *GroupUserStorage) GetGroupsByTutorId(tutorID int64) ([]models.Group, error) {
	var groups []models.Group

	query := `
		SELECT id, title, description, subject_id, tutor_id, tg_group_link, tg_chat_id
		FROM groups
		WHERE tutor_id = $1
	`

	rows, err := g.conn.Query(context.Background(), query, tutorID)
	if err != nil {
		return nil, fmt.Errorf("get all tutor %d groups from db: %w", tutorID, err)
	}

	for rows.Next() {
		var oneGroup models.Group
		err := rows.Scan(
			&oneGroup.ID,
			&oneGroup.Title,
			&oneGroup.Description,
			&oneGroup.SubjectID,
			&oneGroup.TutorID,
			&oneGroup.TgGroupLink,
			&oneGroup.TgChatID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one tutor %d group from db: %w", tutorID, err)
		}

		groups = append(groups, oneGroup)
	}

	return groups, nil
}

func (g *GroupUserStorage) GetGroupUsers(groupID int64) ([]int64, error) {
	var userIDs []int64

	query := `
		SELECT user_id
		FROM user_groups 
		WHERE group_id = $1
	`

	rows, err := g.conn.Query(context.Background(), query, groupID)
	if err != nil {
		return nil, fmt.Errorf("get all group %d users from db: %w", groupID, err)
	}

	for rows.Next() {
		var oneUserID int64

		err := rows.Scan(&oneUserID)
		if err != nil {
			return nil, fmt.Errorf("scan one group %d user from db: %w", groupID, err)
		}

		userIDs = append(userIDs, oneUserID)
	}

	return userIDs, nil
}
