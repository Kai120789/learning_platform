import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { ScheduleResponse } from "../types/types"

export const getScheduleById = createAsyncThunk<
    ScheduleResponse,
    number,
    { rejectValue: string }
>(
    "getScheduleById",
    async (scheduleID, { rejectWithValue }) => {
        try {
            const response = await $api.get<ScheduleResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/${scheduleID}`,
            )

            return response.data
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
