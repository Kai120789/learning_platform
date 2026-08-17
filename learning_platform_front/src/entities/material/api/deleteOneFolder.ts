import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"

export const deleteOneFolder = createAsyncThunk<
    number,
    number,
    { rejectValue: string }
>(
    "deleteOneFolder",
    async (folderID, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/folder/${folderID}`,
            )
            return folderID
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
