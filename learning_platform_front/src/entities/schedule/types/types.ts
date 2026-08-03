export type ScheduleSlotStatus = "FREE" | "BOOKED"

export interface ScheduleSchema {
    data: ScheduleData[] | null
    isLoading: boolean
    error?: string
}

export type ScheduleSlotData = {
    id: number
    scheduleId: number
    startTime: string
    status: ScheduleSlotStatus
    duration?: number | null
    lessonId?: number | null
}

export type ScheduleData = {
    id: number
    tutorId: number
    title?: string | null
    startTime: string
    endTime: string
    slots: ScheduleSlotData[]
}

export type ScheduleSlotResponse = {
    id: number
    schedule_id: number
    start_time: string
    status: ScheduleSlotStatus
    duration?: number | null
    lesson_id?: number | null
}

export type ScheduleResponse = {
    id: number
    tutor_id: number
    title?: string | null
    start_time: string
    end_time: string
    slots: ScheduleSlotResponse[]
}

export type WeekDaySlotView = {
    id: number
    weekday: 0 | 1 | 2 | 3 | 4 | 5 | 6
    date: string
    start: string
    end: string
    title: string
    subtitle?: string
    status: ScheduleSlotStatus
}
