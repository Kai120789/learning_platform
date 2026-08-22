import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { TutorUserData } from "../types/types"

export const addTutorStudents = createAsyncThunk<
    TutorUserData[],
    TutorUserData[],
    { rejectValue: string }
>(
    "addTutorStudents",
    async (students, { rejectWithValue }) => {
        try {
            for (const student of students) {
                await $api.post(
                    `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/student/${student.id}`,
                )
            }
            return students
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
