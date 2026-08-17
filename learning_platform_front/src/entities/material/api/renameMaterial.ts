import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const renameMaterial = createAsyncThunk<
    { materialID: number; title: string },
    { materialID: number; title: string },
    { rejectValue: string }
>(
    "renameMaterial",
    async ({ materialID, title }, { rejectWithValue }) => {
        try {
            await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/${materialID}`,
                null,
                { params: { title } },
            )
            return { materialID, title }
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
