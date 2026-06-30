package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/schedules/internal/dto"
	"learning-platform/schedules/internal/models"
	"learning-platform/schedules/internal/models/enum"
	"time"
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

func (ss *ScheduleSlotsStorage) SetScheduleSlots(
	scheduleID int64,
	slots []dto.CreateScheduleSlot,
) ([]models.ScheduleSlot, error) {
	var resSlots []models.ScheduleSlot
	startTimeSlots := make([]time.Time, 0, len(slots))
	statusSlots := make([]string, 0, len(slots))
	durationSlots := make([]*int64, 0, len(slots))
	lessonIDSlots := make([]*int64, 0, len(slots))
	scheduleIDSlots := make([]int64, 0, len(slots))

	for _, slot := range slots {
		startTimeSlots = append(startTimeSlots, slot.StartTime)
		durationSlots = append(durationSlots, slot.Duration)
		lessonIDSlots = append(lessonIDSlots, slot.LessonID)
		scheduleIDSlots = append(scheduleIDSlots, scheduleID)
		if slot.LessonID != nil {
			statusSlots = append(statusSlots, string(enum.StatusBooked))
		} else {
			statusSlots = append(statusSlots, string(enum.StatusFree))
		}
	}

	query := `
		INSERT INTO schedule_slots (
		    schedule_id,
		    start_time, 
		    status, 
		    duration, 
		    lesson_id
		)
		SELECT *
		FROM unnest(
			$1::bigint[],
			$2::timestamptz[],
			$3::status_enum[],
			$4::bigint[],
		    $5::bigint[]
		)
		RETURNING id, schedule_id, start_time, status, duration, lesson_id
	`

	rows, err := ss.conn.Query(
		context.Background(),
		query,
		scheduleIDSlots,
		startTimeSlots,
		statusSlots,
		durationSlots,
		lessonIDSlots,
	)
	if err != nil {
		return nil, fmt.Errorf("set schedule %d slots: %w", scheduleID, err)
	}

	for rows.Next() {
		var oneSlot models.ScheduleSlot

		err := rows.Scan(
			&oneSlot.ID,
			&oneSlot.ScheduleID,
			&oneSlot.StartTime,
			&oneSlot.Status,
			&oneSlot.Duration,
			&oneSlot.LessonID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one slot for schedule %d: %w", scheduleID, err)
		}

		resSlots = append(resSlots, oneSlot)
	}

	return resSlots, nil
}

func (ss *ScheduleSlotsStorage) UpdateScheduleSlot(
	scheduleSlotID int64,
	scheduleSlot dto.CreateScheduleSlot,
) (*models.ScheduleSlot, error) {
	var resScheduleSlot models.ScheduleSlot
	query := `
		UPDATE schedule_slots
		SET 
		    start_time = $2,
		    duration = $3,
			lesson_id = $4
		WHERE id = $1
		RETURNING id, schedule_id, start_time, status, duration, lesson_id
	`

	err := ss.conn.QueryRow(
		context.Background(),
		query,
		scheduleSlotID,
		scheduleSlot.StartTime,
		scheduleSlot.Duration,
		scheduleSlot.LessonID,
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
