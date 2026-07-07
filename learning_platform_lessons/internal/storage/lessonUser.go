package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LessonUserStorage struct {
	conn *pgxpool.Pool
}

func NewLessonUserStorage(conn *pgxpool.Pool) *LessonUserStorage {
	return &LessonUserStorage{
		conn: conn,
	}
}

func (lu *LessonUserStorage) SetUsersToLesson(lessonID int64, userIDs []int64) error {
	query := `
		INSERT INTO lesson_users (user_id, lesson_id)
		SELECT unnest($1::bigint[]), $2
	`

	_, err := lu.conn.Exec(context.Background(), query, userIDs, lessonID)
	if err != nil {
		return fmt.Errorf("add users %d to lesson %d: %w", userIDs, lessonID, err)
	}

	return nil
}

func (lu *LessonUserStorage) DeleteUsersFromLesson(lessonID int64, userIDs []int64) error {
	query := `
		DELETE FROM lesson_users
		WHERE user_id = ANY($1::bigint[])
		AND lesson_id = $2
	`

	_, err := lu.conn.Exec(context.Background(), query, userIDs, lessonID)
	if err != nil {
		return fmt.Errorf("delete users %d from lesson %d: %w", userIDs, lessonID, err)
	}

	return nil
}

func (lu *LessonUserStorage) GetAllLessonParticipants(lessonID int64) ([]int64, error) {
	var resLessonUserIDs []int64
	query := `
		SELECT user_id
		FROM lesson_users
		WHERE lesson_id = $1
	`

	rows, err := lu.conn.Query(context.Background(), query, lessonID)
	if err != nil {
		return nil, fmt.Errorf("get all users for lesson %d from db: %w", lessonID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneUserID int64

		err := rows.Scan(
			&oneUserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one lesson media from db %w: ", err)
		}

		resLessonUserIDs = append(resLessonUserIDs, oneUserID)
	}

	return resLessonUserIDs, nil
}
