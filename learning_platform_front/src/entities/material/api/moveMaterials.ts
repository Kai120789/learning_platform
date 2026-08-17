import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { MoveMaterialsRequest } from "@/entities/material"

export const moveMaterials = createAsyncThunk<
    MoveMaterialsRequest,
    MoveMaterialsRequest,
    { rejectValue: string }
>(
    "moveMaterials",
    async (request, { rejectWithValue }) => {
        try {
            await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/`,
                request,
            )
            return request
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
