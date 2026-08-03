import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { ScheduleResponse } from "../types/types"

export const getSchedulesByTutorId = createAsyncThunk<
    ScheduleResponse[],
    number,
    { rejectValue: string }
>(
    "getSchedulesByTutorId",
    async (tutorID, { rejectWithValue }) => {
        try {
            const response = await $api.get<ScheduleResponse[] | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/tutor/${tutorID}`,
            )

            return response.data ?? []
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
