import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { GetTutorsRequest, GetTutorsResponse } from "../types/types"

export const getTutors = createAsyncThunk<
    GetTutorsResponse,
    GetTutorsRequest,
    { rejectValue: string }
>(
    "getTutors",
    async ({ page, limit, subjectId }, { rejectWithValue }) => {
        try {
            const response = await $api.get<GetTutorsResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor`,
                {
                    params: {
                        page,
                        limit,
                        ...(subjectId ? { subject_id: subjectId } : {}),
                    },
                },
            )
            return {
                tutors: response.data?.tutors ?? [],
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
