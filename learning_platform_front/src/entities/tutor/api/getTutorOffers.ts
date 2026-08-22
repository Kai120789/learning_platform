import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { TutorOfferResponse } from "../types/types"

export const getTutorOffers = createAsyncThunk<
    TutorOfferResponse[],
    number,
    { rejectValue: string }
>(
    "getTutorOffers",
    async (tutorId, { rejectWithValue }) => {
        try {
            const response = await $api.get<TutorOfferResponse[] | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/offer/${tutorId}`,
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
