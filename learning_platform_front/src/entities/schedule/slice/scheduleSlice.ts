import { createSlice } from "@reduxjs/toolkit"
import type { ScheduleSchema } from "../types/types"
import { mapScheduleResponse, mapScheduleSlotResponse } from "../lib/mapSchedule"
import {
    getSchedulesByTutorId,
    getScheduleById,
    createSchedule,
    createScheduleSlot,
    updateSchedule,
    deleteSchedule,
    deleteScheduleSlot,
    updateScheduleSlot,
    bindLessonToScheduleSlot,
    deleteLessonFromScheduleSlot,
} from "@/entities/schedule"

const initialState: ScheduleSchema = {
    data: null,
    isLoading: false,
    error: undefined,
}

const scheduleSlice = createSlice({
    name: "schedule",
    initialState,
    reducers: {
        clearSchedules: (state) => {
            state.data = null
            state.error = undefined
            state.isLoading = false
        },
    },
    extraReducers: (builder) => {
        builder.addCase(getSchedulesByTutorId.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(getSchedulesByTutorId.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
            if (state.data === null) {
                state.data = []
            }
        })
        builder.addCase(getSchedulesByTutorId.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            state.data = (action.payload ?? []).map(mapScheduleResponse)
        })

        builder.addCase(getScheduleById.pending, (state) => {
            state.isLoading = true
            state.error = ""
        })
        builder.addCase(getScheduleById.rejected, (state, action) => {
            state.isLoading = false
            state.error = action.payload as string
        })
        builder.addCase(getScheduleById.fulfilled, (state, action) => {
            state.isLoading = false
            state.error = ""
            const mapped = mapScheduleResponse(action.payload)
            if (!state.data) {
                state.data = [mapped]
                return
            }
            const index = state.data.findIndex((item) => item.id === mapped.id)
            if (index >= 0) {
                state.data[index] = mapped
            } else {
                state.data.push(mapped)
            }
        })

        builder.addCase(createSchedule.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(createSchedule.fulfilled, (state, action) => {
            state.error = ""
            if (!state.data) {
                state.data = []
            }
            state.data.push(mapScheduleResponse(action.payload))
        })

        builder.addCase(createScheduleSlot.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(createScheduleSlot.fulfilled, (state, action) => {
            state.error = ""
            const mapped = mapScheduleSlotResponse(action.payload)
            const schedule = state.data?.find((item) => item.id === mapped.scheduleId)
            if (!schedule) return
            const index = schedule.slots.findIndex((slot) => slot.id === mapped.id)
            if (index >= 0) {
                schedule.slots[index] = mapped
            } else {
                schedule.slots.push(mapped)
            }
        })

        builder.addCase(updateSchedule.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(updateSchedule.fulfilled, (state, action) => {
            state.error = ""
            const mapped = mapScheduleResponse(action.payload)
            if (!state.data) {
                state.data = [mapped]
                return
            }
            const index = state.data.findIndex((item) => item.id === mapped.id)
            if (index >= 0) {
                const prevSlots = state.data[index].slots
                state.data[index] = {
                    ...mapped,
                    slots: mapped.slots.length > 0 ? mapped.slots : prevSlots,
                }
            } else {
                state.data.push(mapped)
            }
        })

        builder.addCase(deleteSchedule.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(deleteSchedule.fulfilled, (state, action) => {
            state.error = ""
            state.data = state.data?.filter((item) => item.id !== action.meta.arg) ?? []
        })

        builder.addCase(deleteScheduleSlot.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(deleteScheduleSlot.fulfilled, (state, action) => {
            state.error = ""
            for (const schedule of state.data ?? []) {
                const nextSlots = schedule.slots.filter((slot) => slot.id !== action.meta.arg)
                if (nextSlots.length !== schedule.slots.length) {
                    schedule.slots = nextSlots
                    break
                }
            }
        })

        builder.addCase(updateScheduleSlot.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(updateScheduleSlot.fulfilled, (state, action) => {
            state.error = ""
            const mapped = mapScheduleSlotResponse(action.payload)
            const schedule = state.data?.find((item) => item.id === mapped.scheduleId)
            if (!schedule) return
            const index = schedule.slots.findIndex((slot) => slot.id === mapped.id)
            if (index >= 0) {
                schedule.slots[index] = mapped
            } else {
                schedule.slots.push(mapped)
            }
        })

        builder.addCase(bindLessonToScheduleSlot.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(bindLessonToScheduleSlot.fulfilled, (state, action) => {
            state.error = ""
            const { scheduleSlotID, lessonID } = action.meta.arg
            for (const schedule of state.data ?? []) {
                const slot = schedule.slots.find((item) => item.id === scheduleSlotID)
                if (slot) {
                    slot.lessonId = lessonID
                    slot.status = "BOOKED"
                    break
                }
            }
        })

        builder.addCase(deleteLessonFromScheduleSlot.rejected, (state, action) => {
            state.error = action.payload as string
        })
        builder.addCase(deleteLessonFromScheduleSlot.fulfilled, (state, action) => {
            state.error = ""
            for (const schedule of state.data ?? []) {
                const slot = schedule.slots.find((item) => item.id === action.meta.arg)
                if (slot) {
                    slot.lessonId = null
                    slot.status = "FREE"
                    break
                }
            }
        })
    },
})

export const { actions: scheduleActions, reducer: scheduleReducer } = scheduleSlice
