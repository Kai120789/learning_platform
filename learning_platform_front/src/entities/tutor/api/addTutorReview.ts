import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { NewTutorReviewRequest, TutorReviewResponse } from "../types/types"

export const addTutorReview = createAsyncThunk<
    TutorReviewResponse,
    NewTutorReviewRequest,
    { rejectValue: string }
>(
    "addTutorReview",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.post<TutorReviewResponse | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/review`,
                request,
            )
            if (!response.data?.id) {
                return rejectWithValue("Не удалось отправить отзыв")
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
