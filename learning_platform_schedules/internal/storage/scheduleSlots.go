package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models"
	"learning-platform/schedules/internal/models/enum"
)

type ScheduleSlotsStorage struct {
	conn *pgxpool.Pool
}

func NewScheduleSlotsStorage(conn *pgxpool.Pool) *ScheduleSlotsStorage {
	return &ScheduleSlotsStorage{
		conn: conn,
	}
}

func (ss *ScheduleSlotsStorage) GetAllScheduleSlots(scheduleID int64) ([]models.ScheduleSlot, error) {
	var resScheduleSlots []models.ScheduleSlot
	query := `
		SELECT id, schedule_id, start_time, status, duration, lesson_id
		FROM schedule_slots
		WHERE schedule_id = $1
	`

	rows, err := ss.conn.Query(
		context.Background(),
		query,
		scheduleID,
	)
	if err != nil {
		return nil, fmt.Errorf("get all schedule %d slots: %w", scheduleID, err)
	}

	for rows.Next() {
		var oneScheduleSlot models.ScheduleSlot

		err := rows.Scan(
			&oneScheduleSlot.ID,
			&oneScheduleSlot.ScheduleID,
			&oneScheduleSlot.StartTime,
			&oneScheduleSlot.Status,
			&oneScheduleSlot.Duration,
			&oneScheduleSlot.LessonID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one schedule %d slot: %w", scheduleID, err)
		}

		resScheduleSlots = append(resScheduleSlots, oneScheduleSlot)
	}

	return resScheduleSlots, nil
}

func (ss *ScheduleSlotsStorage) CreateScheduleSlot(
	slot dto.CreateScheduleSlot,
) (*models.ScheduleSlot, error) {
	var resSlot models.ScheduleSlot

	query := `
		INSERT INTO schedule_slots (
		    schedule_id,
		    start_time, 
		    status, 
		    duration, 
		    lesson_id
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, schedule_id, start_time, status, duration, lesson_id
	`

	err := ss.conn.QueryRow(
		context.Background(),
		query,
		slot.ScheduleID,
		slot.StartTime,
		string(enum.StatusFree),
		slot.Duration,
		nil,
	).Scan(
		&resSlot.ID,
		&resSlot.ScheduleID,
		&resSlot.StartTime,
		&resSlot.Status,
		&resSlot.Duration,
		&resSlot.LessonID,
	)
	if err != nil {
		return nil, fmt.Errorf("create schedule %d slot: %w", slot.ScheduleID, err)
	}

	return &resSlot, nil
}

func (ss *ScheduleSlotsStorage) UpdateScheduleSlot(
	scheduleSlotID int64,
	scheduleSlot dto.UpdateScheduleSlot,
) (*models.ScheduleSlot, error) {
	var resScheduleSlot models.ScheduleSlot
	query := `
		UPDATE schedule_slots
		SET 
		    start_time = $2,
		    duration = $3
		WHERE id = $1
		RETURNING id, schedule_id, start_time, status, duration, lesson_id
	`

	err := ss.conn.QueryRow(
		context.Background(),
		query,
		scheduleSlotID,
		scheduleSlot.StartTime,
		scheduleSlot.Duration,
	).Scan(
		&resScheduleSlot.ID,
		&resScheduleSlot.ScheduleID,
		&resScheduleSlot.StartTime,
		&resScheduleSlot.Status,
		&resScheduleSlot.Duration,
		&resScheduleSlot.LessonID,
	)
	if err != nil {
		return nil, fmt.Errorf("update schedule slot %d: %w", scheduleSlotID, err)
	}

	return &resScheduleSlot, nil
}

func (ss *ScheduleSlotsStorage) DeleteOneScheduleSlot(scheduleSlotID int64) error {
	query := `
		DELETE FROM schedule_slots
		WHERE id = $1
	`

	_, err := ss.conn.Exec(
		context.Background(),
		query,
		scheduleSlotID,
	)
	if err != nil {
		return fmt.Errorf("delete one schedule slot: %w", err)
	}

	return nil
}

func (ss *ScheduleSlotsStorage) DeleteScheduleSlots(scheduleSlotIDs []int64) error {
	query := `
		DELETE FROM schedule_slots
		WHERE id = ANY($1::bigint[])
	`

	_, err := ss.conn.Exec(
		context.Background(),
		query,
		scheduleSlotIDs,
	)
	if err != nil {
		return fmt.Errorf("delete schedule slots: %w", err)
	}

	return nil
}

func (ss *ScheduleSlotsStorage) DeleteSlotsByScheduleID(scheduleID int64) error {
	query := `
		DELETE FROM schedule_slots
		WHERE schedule_id = $1
	`

	_, err := ss.conn.Exec(
		context.Background(),
		query,
		scheduleID,
	)
	if err != nil {
		return fmt.Errorf("delete slots by schedule id %d: %w", scheduleID, err)
	}

	return nil
}

func (ss *ScheduleSlotsStorage) BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error {
	query := `
		UPDATE schedule_slots
		SET 
		    lesson_id = $2,
		    status = $3
		WHERE id = $1
	`

	_, err := ss.conn.Exec(
		context.Background(),
		query,
		scheduleSlotID,
		lessonID,
		string(enum.StatusBooked),
	)
	if err != nil {
		return fmt.Errorf("bind lesson to schedule slot %d: %w", scheduleSlotID, err)
	}

	return nil
}

func (ss *ScheduleSlotsStorage) DeleteLessonFromScheduleSlot(scheduleSlotID int64) error {
	query := `
		UPDATE schedule_slots
		SET 
		    lesson_id = null,
		    status = $2
		WHERE id = $1
	`

	_, err := ss.conn.Exec(
		context.Background(),
		query,
		scheduleSlotID,
		string(enum.StatusFree),
	)
	if err != nil {
		return fmt.Errorf("delete lesson from schedule slot %d: %w", scheduleSlotID, err)
	}

	return nil
}
