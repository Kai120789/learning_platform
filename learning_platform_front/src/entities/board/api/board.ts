import { $api } from "@/app/providers/storeProvider/config/api"
import { createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"
import type { BoardResponse } from "../model/types"

function rejectMessage(error: unknown) {
    if (axios.isAxiosError(error)) {
        return error.response?.data ?? error.message
    }
    return "Неизвестная ошибка"
}

export const getBoard = createAsyncThunk<
    BoardResponse,
    number,
    { rejectValue: string }
>(
    "getBoard",
    async (boardId, { rejectWithValue }) => {
        try {
            const response = await $api.get<BoardResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/board/${boardId}`,
            )
            return response.data
        } catch (error) {
            return rejectWithValue(rejectMessage(error))
        }
    },
)

export const getTutorBoards = createAsyncThunk<
    BoardResponse[],
    void,
    { rejectValue: string }
>(
    "getTutorBoards",
    async (_, { rejectWithValue }) => {
        try {
            const response = await $api.get<BoardResponse[] | null>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/board`,
            )
            return response.data ?? []
        } catch (error) {
            return rejectWithValue(rejectMessage(error))
        }
    },
)

export const createBoard = createAsyncThunk<
    BoardResponse,
    { title: string },
    { rejectValue: string }
>(
    "createBoard",
    async ({ title }, { rejectWithValue }) => {
        try {
            const response = await $api.post<BoardResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/board`,
                { title },
            )
            return response.data
        } catch (error) {
            return rejectWithValue(rejectMessage(error))
        }
    },
)

export const updateBoard = createAsyncThunk<
    BoardResponse,
    { boardId: number; title: string },
    { rejectValue: string }
>(
    "updateBoard",
    async ({ boardId, title }, { rejectWithValue }) => {
        try {
            const response = await $api.put<BoardResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/board/${boardId}`,
                { title },
            )
            return response.data
        } catch (error) {
            return rejectWithValue(rejectMessage(error))
        }
    },
)

export const deleteBoard = createAsyncThunk<
    number,
    number,
    { rejectValue: string }
>(
    "deleteBoard",
    async (boardId, { rejectWithValue }) => {
        try {
            await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/board/${boardId}`,
            )
            return boardId
        } catch (error) {
            return rejectWithValue(rejectMessage(error))
        }
    },
)
