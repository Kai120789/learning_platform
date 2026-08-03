import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { ScheduleSlotResponse } from "../types/types"
import type { CreateScheduleSlotRequest } from "@/entities/schedule"

export const createScheduleSlot = createAsyncThunk<
    ScheduleSlotResponse,
    CreateScheduleSlotRequest,
    { rejectValue: string }
>(
    "createScheduleSlot",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.post<ScheduleSlotResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/slot/`,
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
