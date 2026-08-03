import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteScheduleSlot = createAsyncThunk<
    void,
    number,
    { rejectValue: string }
>(
    "deleteScheduleSlot",
    async (scheduleSlotID, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/slot/${scheduleSlotID}`,
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
