import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from '@/app/providers/storeProvider/config/api';
import type { UserSettingsRequest, UserSettingsResponse } from '../types/types';
import axios from 'axios';

export const updateUserSettings = createAsyncThunk<
    UserSettingsResponse,
    UserSettingsRequest,
    { rejectValue: string }
>(
    'updateUserSettings',
    async (request, { rejectWithValue }) => {
        try {
            const response = await $api.put<UserSettingsResponse>(
                `${import.meta.env.VITE_SERVER_ENDPOINT}/api/user/settings`,
                request
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