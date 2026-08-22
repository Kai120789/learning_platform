import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { NewTutorOfferRequest, TutorOfferResponse } from "../types/types"

export const addTutorOffer = createAsyncThunk<
    TutorOfferResponse,
    NewTutorOfferRequest,
    { rejectValue: string }
>(
    "addTutorOffer",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.post<TutorOfferResponse | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/tutor/offer`,
                request,
            )
            if (!response.data?.id) {
                return rejectWithValue("Не удалось создать предложение")
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
