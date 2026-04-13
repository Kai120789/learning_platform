package group

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/users/internal/dto"
)

type GroupBaseStorage struct {
	conn *pgxpool.Pool
}

func NewGroupBaseStorage(
	conn *pgxpool.Pool,
) *GroupBaseStorage {
	return &GroupBaseStorage{
		conn: conn,
	}
}

func (g *GroupBaseStorage) CreateGroup(
	groupDto dto.CreateGroup,
) (*int64, error) {
	query := `
		INSERT INTO groups (title, description, subject_id, tutor_id, tg_group_link, tg_chat_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int64
	err := g.conn.QueryRow(
		context.Background(),
		query,
		groupDto.Title,
		groupDto.Description,
		groupDto.SubjectId,
		groupDto.TutorId,
		groupDto.TgGroupLink,
		groupDto.TgChatId,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create group %s in db: %w", groupDto.Title, err)
	}

	return &id, nil
}

func (g *GroupBaseStorage) UpdateGroup(
	id int64,
	groupDto dto.UpdateGroup,
) error {
	query := `
		UPDATE groups
		SET 
		    title = $2, 
		    description = $3,
		    tg_group_link = $4, 
		    tg_chat_id = $5
		WHERE id = $1
	`

	_, err := g.conn.Exec(
		context.Background(),
		query,
		id,
		groupDto.Title,
		groupDto.Description,
		groupDto.TgGroupLink,
		groupDto.TgChatId,
	)
	if err != nil {
		return fmt.Errorf("update group %d in db: %w", id, err)
	}

	return nil
}

func (g *GroupBaseStorage) RemoveGroup(id int64) error {
	query := `DELETE FROM groups WHERE id = $1`

	_, err := g.conn.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("remove group %d in db: %w", id, err)
	}

	return nil
}

func (g *GroupBaseStorage) GetGroupById(id int64) (*models.Group, error) {
	query := `
		SELECT id, title, description, subject_id, tutor_id, tg_group_link, tg_chat_id
		FROM groups
		WHERE id = $1
	`

	row := g.conn.QueryRow(context.Background(), query, id)

	var group models.Group
	err := row.Scan(
		&group.Id,
		&group.Title,
		&group.Description,
		&group.SubjectId,
		&group.TutorId,
		&group.TgGroupLink,
		&group.TgChatId,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("group %d not found in db: %w", id, err)
	}

	if err != nil {
		return nil, fmt.Errorf("get group %d from db: %w", id, err)
	}

	return &group, nil
}

func (g *GroupBaseStorage) GetGroups() ([]models.Group, error) {
	var allGroups []models.Group
	query := `
		SELECT id, title, description, subject_id, tutor_id, tg_group_link, tg_chat_id 
		FROM groups
	`

	rows, err := g.conn.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("get all groups from db: %w", err)
	}

	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.Id,
			&group.Title,
			&group.Description,
			&group.SubjectId,
			&group.TutorId,
			&group.TgGroupLink,
			&group.TgChatId,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one group from db: %w", err)
		}

		allGroups = append(allGroups, group)
	}

	return allGroups, nil
}
