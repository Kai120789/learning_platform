export type CreateScheduleRequest = {
    tutor_id: number
    title: string
    start_time: string
    end_time: string
}

export type CreateScheduleSlotRequest = {
    schedule_id: number
    start_time: string
    duration?: number | null
}
