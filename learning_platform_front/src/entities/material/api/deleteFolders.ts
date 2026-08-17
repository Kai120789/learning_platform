import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteFolders = createAsyncThunk<
    number[],
    number[],
    { rejectValue: string }
>(
    "deleteFolders",
    async (folderIDs, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/folder`,
                { data: { folder_ids: folderIDs } },
            )
            return folderIDs
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
