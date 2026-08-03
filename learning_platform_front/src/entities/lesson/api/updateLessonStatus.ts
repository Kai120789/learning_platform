import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { LessonStatus } from "../types/types"

export const updateLessonStatus = createAsyncThunk<
    { lessonID: number; status: LessonStatus },
    { lessonID: number; status: LessonStatus },
    { rejectValue: string }
>(
    "updateLessonStatus",
    async ({ lessonID, status }, { rejectWithValue }) => {
        try {
            await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/lesson/${lessonID}`,
                JSON.stringify(status),
                {
                    headers: {
                        "Content-Type": "application/json",
                    },
                },
            )
            return { lessonID, status }
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
