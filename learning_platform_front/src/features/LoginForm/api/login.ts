import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios'
import type { LoginRequestDTO } from '..';

export const login = createAsyncThunk(
    'login',
    async (request: LoginRequestDTO, { rejectWithValue }) => {
        try {
            const response = await axios.post(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/auth/login`,
                request,
            )

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || error.message);
        }
    }
)