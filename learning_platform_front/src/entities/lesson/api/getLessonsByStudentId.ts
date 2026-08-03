import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { LessonResponse } from "../types/types"

export const getLessonsByStudentId = createAsyncThunk<
    LessonResponse[],
    number,
    { rejectValue: string }
>(
    "getLessonsByStudentId",
    async (studentID, { rejectWithValue }) => {
        try {
            const response = await $api.get<LessonResponse[] | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/lesson/user/${studentID}`,
            )
            return response.data ?? []
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
