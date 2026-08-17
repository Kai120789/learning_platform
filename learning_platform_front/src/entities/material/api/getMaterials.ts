import { createAsyncThunk } from "@reduxjs/toolkit"
import { $api } from "@/app/providers/storeProvider/config/api"
import axios from "axios"
import type { FoldersAndMaterialsResponse } from "../types/types"

export const getMaterials = createAsyncThunk<
    FoldersAndMaterialsResponse,
    number,
    { rejectValue: string }
>(
    "getMaterials",
    async (folderID, { rejectWithValue }) => {
        try {
            const response = await $api.get<FoldersAndMaterialsResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/material/${folderID}`,
            )
            return {
                folders: response.data?.folders ?? [],
                materials: response.data?.materials ?? [],
            }
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(error.response?.data ?? error.message)
            }
            return rejectWithValue("Неизвестная ошибка")
        }
    }
)
