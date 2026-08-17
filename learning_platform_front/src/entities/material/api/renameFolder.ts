import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const renameFolder = createAsyncThunk<
    { folderID: number; title: string },
    { folderID: number; title: string },
    { rejectValue: string }
>(
    "renameFolder",
    async ({ folderID, title }, { rejectWithValue }) => {
        try {
            await $api.patch(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/folder/${folderID}`,
                null,
                { params: { title } },
            )
            return { folderID, title }
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
