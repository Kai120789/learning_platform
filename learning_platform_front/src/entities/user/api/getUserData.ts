import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import type { UserData } from '../types/types';
import axios from 'axios';

export const getUserData = createAsyncThunk<
    UserData,
    { userId: number },
    { rejectValue: string }
>(
    'getUserData',
    async (userId, { rejectWithValue }) => {
        try {
            const response = await $api.get<UserData>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/user/data/${userId}`,
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