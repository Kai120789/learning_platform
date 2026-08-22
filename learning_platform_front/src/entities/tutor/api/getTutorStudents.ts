import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { GetTutorStudentsRequest, GetTutorStudentsResponse } from "../types/types"

export const getTutorStudents = createAsyncThunk<
    GetTutorStudentsResponse,
    GetTutorStudentsRequest,
    { rejectValue: string }
>(
    "getTutorStudents",
    async ({ page, limit, interactedWithinDays }, { rejectWithValue }) => {
        try {
            const response = await $api.get<GetTutorStudentsResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/student`,
                {
                    params: {
                        page,
                        limit,
                        ...(interactedWithinDays
                            ? { interacted_within_days: interactedWithinDays }
                            : {}),
                    },
                },
            )
            return {
                students: response.data?.students ?? [],
                count: response.data?.count ?? 0,
            }
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
