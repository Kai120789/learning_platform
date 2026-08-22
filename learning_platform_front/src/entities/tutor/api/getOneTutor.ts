import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { OneTutorResponse } from "../types/types"

export type GetOneTutorRequest = {
    tutorId: number
    purpose?: "profile" | "teaching"
}

export const getOneTutor = createAsyncThunk<
    OneTutorResponse,
    GetOneTutorRequest,
    { rejectValue: string }
>(
    "getOneTutor",
    async ({ tutorId }, { rejectWithValue }) => {
        try {
            const response = await $api.get<OneTutorResponse | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/${tutorId}`,
            )
            if (!response.data?.tutor_info?.tutor) {
                return rejectWithValue("Репетитор не найден")
            }
            return response.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
