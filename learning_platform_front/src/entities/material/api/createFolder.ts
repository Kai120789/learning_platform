import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { CreateFolderRequest } from "@/entities/material"
import type { MaterialFolderResponse } from "../types/types"

export const createFolder = createAsyncThunk<
    MaterialFolderResponse,
    CreateFolderRequest,
    { rejectValue: string }
>(
    "createFolder",
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.post<MaterialFolderResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/folder`,
                request,
            )
            return response.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
