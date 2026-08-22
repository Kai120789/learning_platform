import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const updateTutorSubjects = createAsyncThunk<
    number[],
    number[],
    { rejectValue: string }
>(
    "updateTutorSubjects",
    async (subjectIds, { rejectWithValue }) => {
        try {
            const response = await $api.put<number[] | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/subject`,
                subjectIds,
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
