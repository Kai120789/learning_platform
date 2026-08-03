import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { ScheduleResponse } from "../types/types"
import type { CreateScheduleRequest } from "@/entities/schedule"

export const createSchedule = createAsyncThunk<
    ScheduleResponse,
    CreateScheduleRequest,
    { rejectValue: string }
>(
    "createSchedule",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.post<ScheduleResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/`,
                request,
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
