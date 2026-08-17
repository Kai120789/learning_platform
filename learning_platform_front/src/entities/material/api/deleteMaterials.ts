import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteMaterials = createAsyncThunk<
    number[],
    number[],
    { rejectValue: string }
>(
    "deleteMaterials",
    async (materialIDs, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/`,
                { data: { material_ids: materialIDs } },
            )
            return materialIDs
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
