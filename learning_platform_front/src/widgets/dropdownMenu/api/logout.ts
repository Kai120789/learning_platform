import { $api } from '@/app/providers/storeProvider/config/api';
import { createAsyncThunk } from '@reduxjs/toolkit';

export const logout = createAsyncThunk(
    'logout',
    async (__DO_NOT_USE__ActionTypes, { rejectWithValue }) => {
        try {
            const response = await $api.delete(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/auth/logout`,
            )

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || error.message);
        }
    }
)