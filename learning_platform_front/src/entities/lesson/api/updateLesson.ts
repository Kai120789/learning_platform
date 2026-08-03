import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { LessonResponse } from "../types/types"
import type { UpdateLessonRequest } from "../types/updateLesson"

export const updateLesson = createAsyncThunk<
    LessonResponse,
    UpdateLessonRequest,
    { rejectValue: string }
>(
    "updateLesson",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.put<LessonResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/lesson/${request.id}`,
                request,
            )
            return response.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
