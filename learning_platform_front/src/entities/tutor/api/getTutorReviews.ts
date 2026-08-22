import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { GetTutorReviewsRequest, GetTutorReviewsResponse } from "../types/types"

export const getTutorReviews = createAsyncThunk<
    GetTutorReviewsResponse,
    GetTutorReviewsRequest,
    { rejectValue: string }
>(
    "getTutorReviews",
    async ({ tutorId, page, limit }, { rejectWithValue }) => {
        try {
            const response = await $api.get<GetTutorReviewsResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/review/${tutorId}`,
                { params: { page, limit } },
            )
            return {
                reviews: response.data?.reviews ?? [],
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
