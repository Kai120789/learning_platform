import { createAsyncThunk } from '@reduxjs/toolkit';
import type { RegisterRequestDTO, RegisterResponseDTO } from '../types/types';
import { $api } from '@/app/providers/storeProvider/config/api';
import axios from 'axios';

export const register = createAsyncThunk<
    RegisterResponseDTO,
    RegisterRequestDTO,
    { rejectValue: string }
>(
    'register',
    async (request: RegisterRequestDTO, { rejectWithValue }) => {
        try {
            const response = await $api.post<RegisterResponseDTO>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/auth/register`,
                request,
            )

            return response.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                return rejectWithValue(
                    error.response?.data ?? error.message
                );
            }

            return rejectWithValue("Неизвестная ошибка");
        }
    }
)