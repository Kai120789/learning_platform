import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { ScheduleSlotResponse } from "../types/types"
import type { UpdateScheduleSlotRequest } from "@/entities/schedule"

export const updateScheduleSlot = createAsyncThunk<
    ScheduleSlotResponse,
    { scheduleSlotID: number; request: UpdateScheduleSlotRequest },
    { rejectValue: string }
>(
    "updateScheduleSlot",
    async ({ scheduleSlotID, request }, { rejectWithValue }) => {
        try {
            const response = await $api.put<ScheduleSlotResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/schedule/slot/${scheduleSlotID}`,
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
