import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteSchedule = createAsyncThunk<
    void,
    number,
    { rejectValue: string }
>(
    "deleteSchedule",
    async (scheduleID, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/${scheduleID}`,
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
