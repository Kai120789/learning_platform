import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { MoveFoldersRequest } from "@/entities/material"

export const moveFolders = createAsyncThunk<
    MoveFoldersRequest,
    MoveFoldersRequest,
    { rejectValue: string }
>(
    "moveFolders",
    async (request, { rejectWithValue }) => {
        try {
            await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/folder`,
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
