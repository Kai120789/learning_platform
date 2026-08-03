import type { StateSchema } from "@/app/providers/storeProvider"
import type { ScheduleData, ScheduleSlotData } from "../types/types"

export const getSchedules = (state: StateSchema) => state.schedule.data
export const getSchedulesIsLoading = (state: StateSchema) => state.schedule.isLoading
export const getSchedulesError = (state: StateSchema) => state.schedule.error

export const getAllScheduleSlots = (state: StateSchema): ScheduleSlotData[] => {
    return (state.schedule.data ?? []).flatMap((schedule: ScheduleData) => schedule.slots)
}
