import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { MaterialResponse } from "../types/types"

export const createMaterials = createAsyncThunk<
    MaterialResponse[],
    { folderID: number; files: File[] },
    { rejectValue: string }
>(
    "createMaterials",
    async ({ folderID, files }, { rejectWithValue }) => {
        try {
            const formData = new FormData()
            files.forEach((file) => {
                formData.append("files", file)
            })

            const response = await $api.post<MaterialResponse[]>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/${folderID}`,
                formData,
            )
            return response.data ?? []
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
