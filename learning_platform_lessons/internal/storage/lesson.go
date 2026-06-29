package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/lessons/internal/dto"
	"learning-platform/lessons/internal/models"
	"learning-platform/lessons/internal/models/enum"
)

type LessonStorage struct {
	conn *pgxpool.Pool
}

func NewLessonStorage(conn *pgxpool.Pool) *LessonStorage {
	return &LessonStorage{
		conn: conn,
	}
}

func (l *LessonStorage) CreateLesson(lessonDto dto.CreateLesson) (*models.Lesson, error) {
	var resLesson models.Lesson
	query := `
		INSERT INTO lessons (board_id, meet_link, start_time, duration, tutor_id, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, board_id, meet_link, start_time, duration, tutor_id, status
	`

	err := l.conn.QueryRow(
		context.Background(),
		query,
		lessonDto.BoardID,
		lessonDto.MeetLink,
		lessonDto.StartTime,
		lessonDto.Duration,
		lessonDto.TutorID,
		enum.StatusScheduled,
	).Scan(
		&resLesson.ID,
		&resLesson.BoardID,
		&resLesson.MeetLink,
		&resLesson.StartTime,
		&resLesson.Duration,
		&resLesson.TutorID,
		&resLesson.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("insert lesson to db: %w", err)
	}

	return &resLesson, nil
}

func (l *LessonStorage) GetAllLesson() ([]models.Lesson, error) {
	var resLessons []models.Lesson
	query := `
		SELECT id, board_id, meet_link, start_time, duration, tutor_id, status
		FROM lessons
	`

	rows, err := l.conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("get all lessons: %w", err)
	}

	for rows.Next() {
		var oneLesson models.Lesson
		err := rows.Scan(
			&oneLesson.ID,
			&oneLesson.BoardID,
			&oneLesson.MeetLink,
			&oneLesson.StartTime,
			&oneLesson.Duration,
			&oneLesson.TutorID,
			&oneLesson.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one lesson from db: %w", err)
		}

		resLessons = append(resLessons, oneLesson)
	}

	return resLessons, nil
}

func (l *LessonStorage) GetLessonById(lessonID int64) (*models.Lesson, error) {
	var resLesson models.Lesson

	query := `
		SELECT id, board_id, meet_link, start_time, duration, tutor_id, status
		FROM lessons
		WHERE id = $1
	`

	err := l.conn.QueryRow(
		context.Background(),
		query,
		lessonID,
	).Scan(
		&resLesson.ID,
		&resLesson.BoardID,
		&resLesson.MeetLink,
		&resLesson.StartTime,
		&resLesson.Duration,
		&resLesson.TutorID,
		&resLesson.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("get lesson %d by id: %w", lessonID, err)
	}

	return &resLesson, nil
}

func (l *LessonStorage) GetLessonsByUserId(userID int64) ([]models.Lesson, error) {
	var resLessons []models.Lesson
	query := `
		SELECT l.id, l.board_id, l.meet_link, l.start_time, l.duration, l.tutor_id, l.status
		FROM lessons as l
		INNER JOIN lesson_users AS lu
		ON lu.lesson_id = l.id
		WHERE lu.user_id = $1
	`

	rows, err := l.conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("get all user %d lessons: %w", userID, err)
	}

	for rows.Next() {
		var oneLesson models.Lesson
		err := rows.Scan(
			&oneLesson.ID,
			&oneLesson.BoardID,
			&oneLesson.MeetLink,
			&oneLesson.StartTime,
			&oneLesson.Duration,
			&oneLesson.TutorID,
			&oneLesson.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one lesson from db: %w", err)
		}

		resLessons = append(resLessons, oneLesson)
	}

	return resLessons, nil
}

func (l *LessonStorage) GetLessonsByTutorId(tutorID int64) ([]models.Lesson, error) {
	var resLessons []models.Lesson
	query := `
		SELECT id, board_id, meet_link, start_time, duration, tutor_id, status
		FROM lessons
		WHERE tutor_id = $1
	`

	rows, err := l.conn.Query(context.Background(), query, tutorID)
	if err != nil {
		return nil, fmt.Errorf("get all tutor %d lessons: %w", tutorID, err)
	}

	for rows.Next() {
		var oneLesson models.Lesson
		err := rows.Scan(
			&oneLesson.ID,
			&oneLesson.BoardID,
			&oneLesson.MeetLink,
			&oneLesson.StartTime,
			&oneLesson.Duration,
			&oneLesson.TutorID,
			&oneLesson.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one lesson from db: %w", err)
		}

		resLessons = append(resLessons, oneLesson)
	}

	return resLessons, nil
}

func (l *LessonStorage) UpdateLesson(lessonDto dto.UpdateLesson) (*models.Lesson, error) {
	var resLesson models.Lesson
	query := `
		UPDATE lessons 
		SET 
		    board_id = $2, 
		    meet_link = $3, 
		    start_time = $4, 
		    duration = $5
		WHERE id = $1
		RETURNING id, board_id, meet_link, start_time, duration, tutor_id, status
	`

	err := l.conn.QueryRow(
		context.Background(),
		query,
		lessonDto.ID,
		lessonDto.BoardID,
		lessonDto.MeetLink,
		lessonDto.StartTime,
		lessonDto.Duration,
	).Scan(
		&resLesson.ID,
		&resLesson.BoardID,
		&resLesson.MeetLink,
		&resLesson.StartTime,
		&resLesson.Duration,
		&resLesson.TutorID,
		&resLesson.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("insert lesson to db: %w", err)
	}

	return &resLesson, nil
}

func (l *LessonStorage) UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error {
	query := `
		UPDATE lessons
		SET status = $2
		WHERE id = $1
	`

	_, err := l.conn.Exec(context.Background(), query, lessonID, status)
	if err != nil {
		return fmt.Errorf("cancel lesson %d with id: %w", lessonID, err)
	}

	return nil
}
