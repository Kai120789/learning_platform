import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { FoldersAccessRequest } from "@/entities/material"

export const updateUsersFoldersAccess = createAsyncThunk<
    void,
    FoldersAccessRequest,
    { rejectValue: string }
>(
    "updateUsersFoldersAccess",
    async (request, { rejectWithValue }) => {
        try {
            await $api.put(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/folder`,
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
