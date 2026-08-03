import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteLessonFromScheduleSlot = createAsyncThunk<
    void,
    number,
    { rejectValue: string }
>(
    "deleteLessonFromScheduleSlot",
    async (scheduleSlotID, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/slot/${scheduleSlotID}/lesson`,
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
