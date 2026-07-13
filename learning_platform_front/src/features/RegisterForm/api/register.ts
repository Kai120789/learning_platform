import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios'
import type { RegisterRequestDTO } from '../types/types';

export const register = createAsyncThunk(
    'register',
    async (request: RegisterRequestDTO, { rejectWithValue }) => {
        try {
            const response = await axios.post(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/auth/register`,
                request,
            )

            return response.data
        } catch (error) {
            return rejectWithValue(error.response?.data || error.message);
        }
    }
)