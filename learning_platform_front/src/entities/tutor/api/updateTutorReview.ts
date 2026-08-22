import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { TutorReviewResponse, UpdateTutorReviewRequest } from "../types/types"

export const updateTutorReview = createAsyncThunk<
    TutorReviewResponse,
    UpdateTutorReviewRequest,
    { rejectValue: string }
>(
    "updateTutorReview",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.put<TutorReviewResponse | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/review`,
                request,
            )
            if (!response.data?.id) {
                return rejectWithValue("Не удалось обновить отзыв")
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
