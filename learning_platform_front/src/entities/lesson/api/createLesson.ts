import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { LessonResponse } from "../types/types"
import type { CreateLessonRequest } from "@/entities/lesson"

export const createLesson = createAsyncThunk<
    LessonResponse,
    CreateLessonRequest,
    { rejectValue: string }
>(
    "createLesson",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.post<LessonResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/lesson/`,
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
