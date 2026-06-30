package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models"
)

type ScheduleStorage struct {
	conn *pgxpool.Pool
}

func NewScheduleStorage(conn *pgxpool.Pool) *ScheduleStorage {
	return &ScheduleStorage{
		conn: conn,
	}
}

func (s *ScheduleStorage) GetScheduleByID(scheduleID int64) (*models.Schedule, error) {
	var schedule models.Schedule
	query := `
		SELECT id, tutor_id, start_time, end_time
		FROM schedules
		WHERE id = $1
	`

	err := s.conn.QueryRow(
		context.Background(),
		query,
		scheduleID,
	).Scan(
		&schedule.ID,
		&schedule.TutorID,
		&schedule.StartTime,
		&schedule.EndTime,
	)
	if err != nil {
		return nil, fmt.Errorf("get schedule by id %d: %w", scheduleID, err)
	}

	return &schedule, nil
}

func (s *ScheduleStorage) GetAllSchedules() ([]models.Schedule, error) {
	var resSchedules []models.Schedule
	query := `
		SELECT id, tutor_id, start_time, end_time
		FROM schedules
	`

	rows, err := s.conn.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("get all schedules: %w", err)
	}

	for rows.Next() {
		var oneSchedule models.Schedule

		err := rows.Scan(
			&oneSchedule.ID,
			&oneSchedule.TutorID,
			&oneSchedule.StartTime,
			&oneSchedule.EndTime,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one schedule: %w", err)
		}

		resSchedules = append(resSchedules, oneSchedule)
	}

	return resSchedules, nil
}

func (s *ScheduleStorage) GetSchedulesByTutorID(tutorID int64) ([]models.Schedule, error) {
	var resSchedules []models.Schedule
	query := `
		SELECT id, tutor_id, start_time, end_time
		FROM schedules
		WHERE tutor_id = $1
	`

	rows, err := s.conn.Query(
		context.Background(),
		query,
		tutorID,
	)
	if err != nil {
		return nil, fmt.Errorf("get schedules by tutor id %d: %w", tutorID, err)
	}

	for rows.Next() {
		var oneSchedule models.Schedule

		err := rows.Scan(
			&oneSchedule.ID,
			&oneSchedule.TutorID,
			&oneSchedule.StartTime,
			&oneSchedule.EndTime,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one schedule: %w", err)
		}

		resSchedules = append(resSchedules, oneSchedule)
	}

	return resSchedules, nil
}

func (s *ScheduleStorage) CreateSchedule(schedule dto.CreateSchedule) (*models.Schedule, error) {
	var resSchedule models.Schedule
	query := `
		INSERT INTO schedules (tutor_id, start_time, end_time)
		VALUES ($1, $2, $3)
		RETURNING id, tutor_id, start_time, end_time
	`

	err := s.conn.QueryRow(
		context.Background(),
		query,
		schedule.TutorID,
		schedule.StartTime,
		schedule.EndTime,
	).Scan(
		&resSchedule.ID,
		&resSchedule.TutorID,
		&resSchedule.StartTime,
		&resSchedule.EndTime,
	)
	if err != nil {
		return nil, fmt.Errorf("create schedule: %w", err)
	}

	return &resSchedule, nil
}

func (s *ScheduleStorage) UpdateSchedule(schedule dto.UpdateSchedule) (*models.Schedule, error) {
	var resSchedule models.Schedule
	query := `
		UPDATE schedules
		SET start_time = $2, end_time = $3
		WHERE id = $1
		RETURNING id, tutor_id, start_time, end_time
	`

	err := s.conn.QueryRow(
		context.Background(),
		query,
		schedule.ID,
		schedule.StartTime,
		schedule.EndTime,
	).Scan(
		&resSchedule.ID,
		&resSchedule.TutorID,
		&resSchedule.StartTime,
		&resSchedule.EndTime,
	)
	if err != nil {
		return nil, fmt.Errorf("update schedule: %w", err)
	}

	return &resSchedule, nil
}

func (s *ScheduleStorage) DeleteSchedule(scheduleID int64) error {
	query := `
		DELETE FROM schedules
		WHERE id = $1
	`

	_, err := s.conn.Exec(
		context.Background(),
		query,
		scheduleID,
	)
	if err != nil {
		return fmt.Errorf("delete schedule %d: %w", scheduleID, err)
	}

	return nil
}
