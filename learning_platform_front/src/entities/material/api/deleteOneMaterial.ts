import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteOneMaterial = createAsyncThunk<
    number,
    number,
    { rejectValue: string }
>(
    "deleteOneMaterial",
    async (materialID, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/${materialID}`,
            )
            return materialID
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
