import { createAsyncThunk } from '@reduxjs/toolkit';
import type { LoginRequestDTO } from '..';
import { $api } from '@/app/providers/storeProvider/config/api';

export const login = createAsyncThunk(
    'login',
    async (request: LoginRequestDTO, { rejectWithValue }) => {
        try {
            const response = await $api.post(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/auth/login`,
                request,
            )

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || error.message);
        }
    }
)