import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { ScheduleResponse } from "../types/types"
import type { UpdateScheduleRequest } from "@/entities/schedule"

export const updateSchedule = createAsyncThunk<
    ScheduleResponse,
    { scheduleID: number; request: UpdateScheduleRequest },
    { rejectValue: string }
>(
    "updateSchedule",
    async ({ scheduleID, request }, { rejectWithValue }) => {
        try {
            const response = await $api.put<ScheduleResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/${scheduleID}`,
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
