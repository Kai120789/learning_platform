import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const bindLessonToScheduleSlot = createAsyncThunk<
    void,
    { scheduleSlotID: number; lessonID: number },
    { rejectValue: string }
>(
    "bindLessonToScheduleSlot",
    async ({ scheduleSlotID, lessonID }, { rejectWithValue }) => {
        try {
            await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/slot/${scheduleSlotID}`,
                lessonID,
            )
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(
                    error.response?.data ?? error.message
                )
            }

            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
