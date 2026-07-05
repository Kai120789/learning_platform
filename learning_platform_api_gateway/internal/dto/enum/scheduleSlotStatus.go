package enum

type ScheduleSlotStatus string

const (
	StatusFree   ScheduleSlotStatus = "FREE"
	StatusBooked ScheduleSlotStatus = "BOOKED"
)
