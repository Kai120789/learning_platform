import type { ScheduleData, ScheduleResponse, ScheduleSlotData, ScheduleSlotResponse } from "../types/types"

export function mapScheduleSlotResponse(slot: ScheduleSlotResponse): ScheduleSlotData {
    return {
        id: slot.id,
        scheduleId: slot.schedule_id,
        startTime: slot.start_time,
        status: slot.status,
        duration: slot.duration ?? null,
        lessonId: slot.lesson_id ?? null,
    }
}

export function mapScheduleResponse(schedule: ScheduleResponse): ScheduleData {
    return {
        id: schedule.id,
        tutorId: schedule.tutor_id,
        title: schedule.title ?? null,
        startTime: schedule.start_time,
        endTime: schedule.end_time,
        slots: (schedule.slots ?? []).map(mapScheduleSlotResponse),
    }
}
