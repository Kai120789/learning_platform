import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { MaterialsAccessRequest } from "@/entities/material"

export const updateUsersMaterialsAccess = createAsyncThunk<
    void,
    MaterialsAccessRequest,
    { rejectValue: string }
>(
    "updateUsersMaterialsAccess",
    async (request, { rejectWithValue }) => {
        try {
            await $api.put(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/`,
                request,
            )
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
