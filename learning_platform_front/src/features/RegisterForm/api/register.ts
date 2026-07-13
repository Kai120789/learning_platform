import { createAsyncThunk } from '@reduxjs/toolkit';
import type { RegisterRequestDTO } from '../types/types';
import { $api } from '@/app/providers/storeProvider/config/api';

export const register = createAsyncThunk(
    'register',
    async (request: RegisterRequestDTO, { rejectWithValue }) => {
        try {
            const response = await $api.post(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/auth/register`,
                request,
            )

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || error.message);
        }
    }
)